package message

import (
	"encoding/json"
	"log"
	"unicode/utf16"

	"github.com/NicoNex/echotron/v3"
)

// UpdateMessage is the custom type for incoming or just sent message (of any
// type). It implements editable interface. Tips: Use the json- tags naming to
// refer to the official Telegram documentation available at:
// https://core.telegram.org/bots/api#message In the rare cases where the json
// string start with "parrbot_" then is not a copy-paste of the response
type UpdateMessage struct {
	// Normal Telegram / Echotron fields
	ID              int                            `json:"message_id"`
	From            *echotron.User                 `json:"from,omitempty"` // <nil> if channel
	SenderChat      *echotron.Chat                 `json:"sender_chat,omitempty"`
	Date            int                            `json:"date"`
	Chat            *echotron.Chat                 `json:"chat"`
	EditDate        int                            `json:"edit_date,omitempty"`
	AuthorSignature string                         `json:"author_signature,omitempty"`
	InlineKeyboard  *echotron.InlineKeyboardMarkup `json:"reply_markup,omitempty"` // it Changed the name: ReplyMarkup (too generic) -> InlineKeyboard
	ViaBot          *echotron.User                 `json:"via_bot,omitempty"`

	ReplyToMessage *UpdateMessage `json:"parrbot_reply_to_message,omitempty"`

	/* Custom wrappers of information about a specific incoming message type
	 * tips: Thanks to that if you want to check if a message for example
	 *       contains a media you can just check if message.Media != nil
	 */
	Forward            *ForwardInfo            `json:"parrbot_forward,omitempty"`
	Media              *MediaInfo              `json:"parrbot_media,omitempty"`
	SystemNotification *SystemNotificationInfo `json:"parrbot_system_notification,omitempty"`

	/* They countain normal text information or the onse of media caption
	 * when text is empty. Notice however how inside Media (of type *MediaInfo)
	 * there is still a capy of the media caption
	 */
	Text     string                    `json:"text,omitempty"`
	Entities []*echotron.MessageEntity `json:"entities,omitempty"`
}

// castMessage transform an *echotron.Message into a *UpdateMessage
func castMessage(original *echotron.Message) (message *UpdateMessage) {
	if original == nil { // Guard close
		return nil
	}

	// Util function for error checking
	check := func(e error) {
		if e != nil {
			log.Fatal(e)
		}
	}

	// Get JSON format of the original
	var jsonData, err = json.Marshal(*original)
	check(err)

	// Copy common values to the new message...
	message = new(UpdateMessage)
	check(json.Unmarshal(jsonData, message))
	if original.ReplyToMessage != nil {
		message.ReplyToMessage = castMessage(original.ReplyToMessage)
	}

	// ... values if the message is forwarded
	forwardMsg := ForwardInfo{}
	check(json.Unmarshal(jsonData, &forwardMsg))
	if b, _ := json.Marshal(forwardMsg); len(b) > 2 {
		message.Forward = &forwardMsg
	}

	// ... values if the message contains media or special attachments
	mediaMsg := MediaInfo{}
	check(json.Unmarshal(jsonData, &mediaMsg))
	if b, _ := json.Marshal(mediaMsg); len(b) > 2 {
		message.Media = &mediaMsg
	}

	// ... values if the message is a Telegram's event message
	SystemMsg := SystemNotificationInfo{}
	check(json.Unmarshal(jsonData, &SystemMsg))
	if original.PinnedMessage != nil {
		SystemMsg.PinnedMessage = castMessage(original.PinnedMessage)
		message.SystemNotification = &SystemMsg
	} else if b, _ := json.Marshal(SystemMsg); len(b) > 2 {
		message.SystemNotification = &SystemMsg
	}

	// Get media the caption if text is an empty string
	if message.Text == "" {
		message.Text = original.Caption
		message.Entities = original.CaptionEntities
	}

	return
}

// EditText is a method that allows to edit the text (and others options)
// for textual messages (message.Text) sent by the bot
func (message *UpdateMessage) EditText(text string, opts *echotron.MessageTextOptions) error {
	return editText(message, text, opts)
}

// EditMedia is a method that allows to edit the media (and others options)
// ONLY for messages sent by the bot that contain it
func (message *UpdateMessage) EditMedia(media echotron.InputMedia, opts *echotron.MessageReplyMarkup) error {
	return editMedia(message, media, opts)
}

// EditInlineKeyboard is a method that allows to edit the InlineKeyboard ONLY
// for messages sent by the bot
func (message *UpdateMessage) EditInlineKeyboard(keyboard [][]echotron.InlineKeyboardButton) error {
	return editInlineKbd(message, keyboard)
}

// EditLiveLocation is a method that allows to edit the Location (and others options)
// ONLY for messages sent by the bot that contain it (like Photo or Document...)
func (message *UpdateMessage) EditLiveLocation(latitude, longitude float64, opts *echotron.EditLocationOptions) error {
	return editLiveLocation(message, latitude, longitude, opts)
}

// EditCaption is a method that allows to edit the caption (and others options)
// ONLY for messages sent by the bot that contain media (like Photo or Document...)
func (message *UpdateMessage) EditCaption(opts *echotron.MessageCaptionOptions) error {
	return editCaption(message, opts)
}

// Delete the given message on the original chat (given message will not sync)
func (message UpdateMessage) Delete() error {
	return delete(&message)
}

/* --- Handle echotron.MessageEntity --- */

// EntityFilter is a function checks an entity, returning true if pass false otherwise
type EntityFilter func(entity echotron.MessageEntity) bool

// ExtractEntitiesOfType extracts the content of the Entities contained filtered by their type
func (message UpdateMessage) ExtractEntitiesOfType(allowedTypes ...echotron.MessageEntityType) []string {
	return message.ExtractEntities(FilterEntityByType(allowedTypes...))
}

// ExtractEntities extracts the content of the Entities contained inside given
// message and filtered by given filter function
func (message UpdateMessage) ExtractEntities(filter EntityFilter) (extracted []string) {
	var text = stirngToUft16(message.Text)

	for _, entity := range message.Entities {
		if entity != nil && filter(*entity) {
			extracted = append(extracted, grabEntityContent(text, *entity))
		}
	}
	return
}

// FilterEntityByType generates an EntityFilter that will return true only if
// entity's type match at least one of the given ones
func FilterEntityByType(allowedTypes ...echotron.MessageEntityType) EntityFilter {
	return func(entity echotron.MessageEntity) bool {
		return contains(allowedTypes, entity.Type)
	}
}

// stirngToUft16 encode a given string into Utf16 format
func stirngToUft16(toEncode string) []uint16 {
	return utf16.Encode([]rune(toEncode))
}

// uft16ToStirng decode a given encoded text in Utf16 format into a string
func uft16ToStirng(encodedText []uint16) string {
	return string(utf16.Decode(encodedText))
}

// grabEntityContent extract the entity content from a given utf16-encoded text
func grabEntityContent(uft16Text []uint16, entity echotron.MessageEntity) string {
	return uft16ToStirng(uft16Text[entity.Offset : entity.Offset+entity.Length])
}

// contains checks if the value of given elem is present on the given list
func contains[T comparable](list []T, elem T) bool {
	for _, value := range list {
		if value == elem {
			return true
		}
	}
	return false
}
