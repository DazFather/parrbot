package message

import (
	"encoding/json"
	"log"

	"github.com/NicoNex/echotron/v3"
)

// Parr(b)ot Update type - it have UpdateMessage instead of echotron.Message
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

// Parr(b)ot Update type - it have UpdateMessage instead of echotron.Message
type CallbackQuery struct {
	ID              string         `json:"id"`
	From            *echotron.User `json:"from"`
	Message         *UpdateMessage `json:"parrbot_message,omitempty"`
	InlineMessageID string         `json:"inline_message_id,omitempty"`
	ChatInstance    string         `json:"chat_instance,omitempty"`
	Data            string         `json:"data,omitempty"`
	GameShortName   string         `json:"game_short_name,omitempty"`
}

// Incoming update types
type UpdateType uint16

const (
	MESSAGE              UpdateType = 1   // 0000000001
	EDITED_MESSAGE                  = 2   // 0000000010
	CHANNEL_POST                    = 4   // 0000000100
	EDITED_CHANNEL_POST             = 8   // 0000001000
	INLINE_QUERY                    = 16  // 0000010000
	CHOSEN_INLINE_RESULT            = 32  // 0000100000
	CALLBACK_QUERY                  = 64  // 0001000000
	MY_CHAT_MEMBER                  = 128 // 0010000000
	CHAT_MEMBER                     = 256 // 0100000000
	CHAT_JOIN_REQUEST               = 512 // 1000000000
)
const ANY UpdateType = 1023 // 1111111111

/* Incoming or just sent message (of any type)
 * tips: Use the json naming to refer to the official Telegram documentation
 *       avaiable at: https://core.telegram.org/bots/api#message
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
	 *       countain a media you can jsut check if message.Media != nil
	 */
	Forward            *ForwardInfo            `json:"parrbot_forward,omitempty"`
	Media              *MediaInfo              `json:"parrbot_media,omitempty"`
	SystemNotification *SystemNotificationInfo `json:"parrbot_system_notification,omitempty"`

	/* They countain normal text information or the onse of  media caption
	 * when text is empty. Notice however how inside Media (of type *MediaInfo)
	 * there is still a capy of the media caption
	 */
	Text     string                    `json:"text,omitempty"`
	Entities []*echotron.MessageEntity `json:"entities,omitempty"`
}

// Infos about the original message that has been forwarded (part of UpdateMessage)
type ForwardInfo struct {
	From       *echotron.User `json:"forward_from,omitempty"`
	Chat       *echotron.Chat `json:"forward_from_chat,omitempty"`
	MessageID  int            `json:"forward_from_message_id,omitempty"`
	Signature  string         `json:"forward_signature,omitempty"`
	SenderName string         `json:"forward_sender_name,omitempty"`
	Date       int            `json:"forward_date,omitempty"`
}

// Infos about the media contained into a message (part of UpdateMessage)
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

// Information about system message that are generated by Telegram after a particular event (part of UpdateMessage)
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

func castMessage(original *echotron.Message) (message *UpdateMessage) {
	if original == nil {
		return nil
	}

	check := func(e error) {
		if e != nil {
			log.Fatal(e)
		}
	}

	var jsonData, err = json.Marshal(*original)
	check(err)

	message = new(UpdateMessage)
	check(json.Unmarshal(jsonData, message))
	if original.ReplyToMessage != nil {
		message.ReplyToMessage = castMessage(original.ReplyToMessage)
	}

	forwardMsg := ForwardInfo{}
	check(json.Unmarshal(jsonData, &forwardMsg))
	if b, _ := json.Marshal(forwardMsg); len(b) > 2 {
		message.Forward = &forwardMsg
	}

	mediaMsg := MediaInfo{}
	check(json.Unmarshal(jsonData, &mediaMsg))
	if b, _ := json.Marshal(mediaMsg); len(b) > 2 {
		message.Media = &mediaMsg
	}

	SystemMsg := SystemNotificationInfo{}
	check(json.Unmarshal(jsonData, &SystemMsg))
	if original.PinnedMessage != nil {
		SystemMsg.PinnedMessage = castMessage(original.PinnedMessage)
		message.SystemNotification = &SystemMsg
	} else if b, _ := json.Marshal(SystemMsg); len(b) > 2 {
		message.SystemNotification = &SystemMsg
	}

	if message.Text == "" {
		message.Text = original.Caption
		message.Entities = original.CaptionEntities
	}

	return
}

func castCallbackQuery(original *echotron.CallbackQuery) (callback *CallbackQuery) {
	if original == nil {
		return nil
	}

	check := func(e error) {
		if e != nil {
			log.Fatal(e)
		}
	}

	var jsonData, err = json.Marshal(*original)
	check(err)

	callback = new(CallbackQuery)
	check(json.Unmarshal(jsonData, callback))

	if original.Message != nil {
		callback.Message = castMessage(original.Message)
	}

	return
}

// Transform *echotron.Update into *Update
func CastUpdate(original *echotron.Update) (update *Update) {
	if original == nil {
		return nil
	}

	check := func(e error) {
		if e != nil {
			log.Fatal(e)
		}
	}

	var jsonData, err = json.Marshal(*original)
	check(err)

	update = new(Update)
	check(json.Unmarshal(jsonData, update))

	update.Message = castMessage(original.Message)
	update.EditedMessage = castMessage(original.EditedMessage)
	update.ChannelPost = castMessage(original.ChannelPost)
	update.EditedChannelPost = castMessage(original.EditedChannelPost)

	update.CallbackQuery = castCallbackQuery(original.CallbackQuery)

	return
}
