package message

import (
	"encoding/json"
	"log"

	"github.com/NicoNex/echotron/v3"
)

/* Update is the Parr(b)ot rapresentation the echotron.Update, the difference
 * is that it have the custom UpdateMessage instead of echotron.Message type.
 * You can cast a echotron.Update to a Update using the CastUpdate function
 */
type Update struct {
	ID                 int                          `json:"update_id"`
	Message            *UpdateMessage               `json:"parrbot_message,omitempty"`
	EditedMessage      *UpdateMessage               `json:"parrbot_edited_message,omitempty"`
	ChannelPost        *UpdateMessage               `json:"parrbot_channel_post,omitempty"`
	EditedChannelPost  *UpdateMessage               `json:"parrbot_edited_channel_post,omitempty"`
	InlineQuery        *echotron.InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *echotron.ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery               `json:"parrbot_callback_query,omitempty"`
	MyChatMember       *echotron.ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *echotron.ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *echotron.ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

/* CallbackQuery is the Parr(b)ot rapresentation the echotron.CallbackQuery,
 * the difference is that it have the custom UpdateMessage instead of echotron.Message type
 */
type CallbackQuery struct {
	ID              string         `json:"id"`
	From            *echotron.User `json:"from"`
	Message         *UpdateMessage `json:"parrbot_message,omitempty"`
	InlineMessageID string         `json:"inline_message_id,omitempty"`
	ChatInstance    string         `json:"chat_instance,omitempty"`
	Data            string         `json:"data,omitempty"`
	GameShortName   string         `json:"game_short_name,omitempty"`
}

// UpdateType represent a possible incoming Update types used on the "ReplyAt" Command inside the command list
type UpdateType uint16

/* These are all the possible types of Update. On the side the binary representation
 * Each one can be used as a flag into the "ReplyAt" field of a Command on the
 * command list
 * tips: You can even sum them to specify that the command will be executed on
 *       reply at for example MESSAGE + CHANNEL_POST (normal written messages)
 *       and channel posts. If you want all, you can use ANY
 */
const (
	MESSAGE              UpdateType = 1 << iota         // 0000000001
	EDITED_MESSAGE                                      // 0000000010
	CHANNEL_POST                                        // 0000000100
	EDITED_CHANNEL_POST                                 // 0000001000
	INLINE_QUERY                                        // 0000010000
	CHOSEN_INLINE_RESULT                                // 0000100000
	CALLBACK_QUERY                                      // 0001000000
	MY_CHAT_MEMBER                                      // 0010000000
	CHAT_MEMBER                                         // 0100000000
	CHAT_JOIN_REQUEST                                   // 1000000000

	// ANY represents any possible UpdateType.
	ANY                             = (1 << iota) - 1   // 1111111111
)

/* UpdateMessage is the custom type for incoming or just sent message (of any type)
 * tips: Use the json naming to refer to the official Telegram documentation
 *       available at: https://core.telegram.org/bots/api#message
 *       In the rare cases where the json string start with "parrbot_" then
 *       is not a copy-paste of the response
 */
type UpdateMessage struct {
	// Normal Telegram / Echotron resonse
	ID              int                            `json:"message_id"`
	Date            int                            `json:"date"`
	From            *echotron.User                 `json:"from,omitempty"` // <nil> if channel
	Chat            *echotron.Chat                 `json:"chat"`
	SenderChat      *echotron.Chat                 `json:"sender_chat,omitempty"`
	EditDate        int                            `json:"edit_date,omitempty"`
	AuthorSignature string                         `json:"author_signature,omitempty"`
	InlineKeyboard  *echotron.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ViaBot          *echotron.User                 `json:"via_bot,omitempty"`

	ReplyToMessage *UpdateMessage `json:"parrbot_reply_to_message,omitempty"`

	/* Custom wrappers of information about a specific incoming message type
	 * tips: Thanks to that if you want to check if a message for example
	 *       countains a media you can just check if message.Media != nil
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

// ForwardInfo countain all the infos of the original message that has been forwarded
type ForwardInfo struct {
	From       *echotron.User `json:"forward_from,omitempty"`
	Chat       *echotron.Chat `json:"forward_from_chat,omitempty"`
	MessageID  int            `json:"forward_from_message_id,omitempty"`
	Signature  string         `json:"forward_signature,omitempty"`
	SenderName string         `json:"forward_sender_name,omitempty"`
	Date       int            `json:"forward_date,omitempty"`
}

// MediaInfo countain all the infos about medias, Polls and so on contained into a message
type MediaInfo struct {
	MediaGroupID    string                    `json:"media_group_id,omitempty"`
	Animation       *echotron.Animation       `json:"animation,omitempty"`
	Audio           *echotron.Audio           `json:"audio,omitempty"`
	Document        *echotron.Document        `json:"document,omitempty"`
	Photo           []*echotron.PhotoSize     `json:"photo,omitempty"`
	Sticker         *echotron.Sticker         `json:"sticker,omitempty"`
	Video           *echotron.Video           `json:"video,omitempty"`
	VideoNote       *echotron.VideoNote       `json:"video_note,omitempty"`
	Voice           *echotron.Voice           `json:"voice,omitempty"`
	Caption         string                    `json:"caption,omitempty"`
	CaptionEntities []*echotron.MessageEntity `json:"caption_entities,omitempty"`
	Contact         *echotron.Contact         `json:"contact,omitempty"`
	Dice            *echotron.Dice            `json:"dice,omitempty"`
	Game            *echotron.Game            `json:"game,omitempty"`
	Poll            *echotron.Poll            `json:"poll,omitempty"`
	Venue           *echotron.Venue           `json:"venue,omitempty"`
	Location        *echotron.Location        `json:"location,omitempty"`
}

// SystemNotificationInfo countain the infos of Telegram's generated message on particular events (part of UpdateMessage)
type SystemNotificationInfo struct {
	NewChatMembers                []*echotron.User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *echotron.User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                                  `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*echotron.PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                                    `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                                    `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                                    `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                                    `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *echotron.MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int                                     `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int                                     `json:"migrate_from_chat_id,omitempty"`
	ConnectedWebsite              string                                  `json:"connected_website,omitempty"`
	ProximityAlertTriggered       *echotron.ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatStarted              *echotron.VoiceChatStarted              `json:"voice_chat_started,omitempty"`
	VoiceChatEnded                *echotron.VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  *echotron.VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	PinnedMessage                 *UpdateMessage                          `json:"parrbot_pinned_message,omitempty"`
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

	// ... values if the message countains media or special attachments
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

// castCallbackQuery transform an *echotron.CallbackQuery into a *CallbackQuery
func castCallbackQuery(original *echotron.CallbackQuery) (callback *CallbackQuery) {
	if original == nil { // Guard close
		return nil
	}

	// Get JSON format of the original
	var jsonData, err = json.Marshal(*original)
	if err != nil {
		log.Fatal(err)
	}

	// Copy common values to the new callback
	callback = new(CallbackQuery)
	if err = json.Unmarshal(jsonData, callback); err != nil {
		log.Fatal(err)
	}

	// Cast *echotron.Message into *UpdateMessage
	if original.Message != nil {
		callback.Message = castMessage(original.Message)
	}

	return
}

// CastUpdate transform an *echotron.Update into a *Update
func CastUpdate(original *echotron.Update) (update *Update) {
	if original == nil { // Guard close
		return nil
	}

	// Get JSON format of the original echotron.Update
	var jsonData, err = json.Marshal(*original)
	if err != nil {
		log.Fatal(err)
	}

	// Copy common values to the new update
	update = new(Update)
	if err = json.Unmarshal(jsonData, update); err != nil {
		log.Fatal(err)
	}

	// Cast *echotron.Message into *UpdateMessage
	update.Message = castMessage(original.Message)
	update.EditedMessage = castMessage(original.EditedMessage)
	update.ChannelPost = castMessage(original.ChannelPost)
	update.EditedChannelPost = castMessage(original.EditedChannelPost)

	// Cast *echotron.CallbackQuery into *CallbackQuery
	update.CallbackQuery = castCallbackQuery(original.CallbackQuery)

	return
}
