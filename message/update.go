package message

import (
	"encoding/json"
	"log"

	"github.com/NicoNex/echotron/v3"
)

// Update is the Parr(b)ot rapresentation the echotron.Update, the difference
// is that it have the custom UpdateMessage (declared on inside the "incoming.go"
// file of the same packages) instead of echotron.Message type.
// You can cast a echotron.Update to a Update using the CastUpdate function
// Update implements editable interface
type Update struct {
	ID                 int                          `json:"update_id"`
	Message            *UpdateMessage               `json:"parrbot_message,omitempty"`
	EditedMessage      *UpdateMessage               `json:"parrbot_edited_message,omitempty"`
	ChannelPost        *UpdateMessage               `json:"parrbot_channel_post,omitempty"`
	EditedChannelPost  *UpdateMessage               `json:"parrbot_edited_channel_post,omitempty"`
	InlineQuery        *echotron.InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *echotron.ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery               `json:"parrbot_callback_query,omitempty"`
	ShippingQuery      *echotron.ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *echotron.PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	MyChatMember       *echotron.ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         *echotron.ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    *echotron.ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

// UpdateType represent a possible incoming Update types used on the "ReplyAt" Command inside the command list
type UpdateType uint16

// These are all the possible types of Update. On the side the binary representation
// Each one can be used as a flag into the "ReplyAt" field of a Command on the
// command list. Tips: You can even sum them to specify that the command will be
// executed onreply at for example MESSAGE + CHANNEL_POST (normal written messages
// and channel posts. If you want all, you can use ANY
const (
	MESSAGE              UpdateType = 1 << iota // 000000000001
	EDITED_MESSAGE                              // 000000000010
	CHANNEL_POST                                // 000000000100
	EDITED_CHANNEL_POST                         // 000000001000
	INLINE_QUERY                                // 000000010000
	CHOSEN_INLINE_RESULT                        // 000000100000
	CALLBACK_QUERY                              // 000001000000
	SHIPPING_QUERY                              // 000010000000
	PRE_CHECKOUT_QUERY                          // 000100000000
	MY_CHAT_MEMBER                              // 001000000000
	CHAT_MEMBER                                 // 010000000000
	CHAT_JOIN_REQUEST                           // 100000000000

	// ANY represents any possible UpdateType.
	ANY = (1 << iota) - 1 // 1111111111
)

// ForwardInfo countain all the infos of the original message that has been forwarded
type ForwardInfo struct {
	From       *echotron.User `json:"forward_from,omitempty"`
	Chat       *echotron.Chat `json:"forward_from_chat,omitempty"`
	MessageID  int            `json:"forward_from_message_id,omitempty"`
	Signature  string         `json:"forward_signature,omitempty"`
	SenderName string         `json:"forward_sender_name,omitempty"`
	Date       int            `json:"forward_date,omitempty"`
	Automatic  bool           `json:"is_automatic_forward_date,omitempty"`
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
	VideoChatScheduled            *echotron.VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *echotron.VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *echotron.VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *echotron.VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *echotron.WebAppData                    `json:"web_app_data,omitempty"`
	PinnedMessage                 *UpdateMessage                          `json:"parrbot_pinned_message,omitempty"`
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

// FromMessage gets the original message contain in the update if present
func (u Update) FromMessage() (msg *UpdateMessage) {
	return u.grabMessage()
}

// Deletes the original message contain in the update if present
func (u Update) DeleteMessage() error {
	return delete(u)
}
