package message

import (
	"github.com/NicoNex/echotron/v3"
)

// CallbackQuery is the Parr(b)ot rapresentation the echotron.CallbackQuery,
// the difference is that it have the custom *UpdateMessage instead of *echotron.Message
// type for Message field. CallbackQuery implements editable interface
type CallbackQuery struct {
	ID              string         `json:"id"`
	From            *echotron.User `json:"from"`
	Message         *UpdateMessage `json:"parrbot_message,omitempty"`
	InlineMessageID string         `json:"inline_message_id,omitempty"`
	ChatInstance    string         `json:"chat_instance,omitempty"`
	Data            string         `json:"data,omitempty"`
	GameShortName   string         `json:"game_short_name,omitempty"`
}

// AnswerAlert allows to reply to a given callback with a _ notification.
// The notification may be cached client-side by final user for cacheTime seconds
// max: 3600s
func (callback CallbackQuery) AnswerAlert(text string, cacheTime uint16) error {
	return callback.Answer(&echotron.CallbackQueryOptions{
		Text:      text,
		CacheTime: int(cacheTime),
		ShowAlert: true,
	})
}

// AnswerToast allows to reply to a given callback with a _ notification.
// The notification may be cached client-side by final user for cacheTime seconds
// max: 3600s
func (callback CallbackQuery) AnswerToast(text string, cacheTime uint16) error {
	return callback.Answer(&echotron.CallbackQueryOptions{
		Text:      text,
		CacheTime: int(cacheTime),
	})
}

// Answer allows to reply to a given callback using given options
func (callback CallbackQuery) Answer(opts *echotron.CallbackQueryOptions) error {
	return parseResponseError(api.AnswerCallbackQuery(callback.ID, opts))
}

// EditText is a method that allows to edit the text (and others options)
// for textual messages (message.Text) sent by the bot
func (callback CallbackQuery) EditText(text string, opts *echotron.MessageTextOptions) error {
	return editText(callback, text, opts)
}

// EditMedia is a method that allows to edit the media (and others options)
// ONLY for messages sent by the bot that contain it
func (callback CallbackQuery) EditMedia(media echotron.InputMedia, opts *echotron.MessageReplyMarkup) error {
	return editMedia(callback, media, opts)
}

// EditInlineKeyboard is a method that allows to edit the InlineKeyboard ONLY
// for messages sent by the bot
func (callback CallbackQuery) EditInlineKeyboard(keyboard [][]echotron.InlineKeyboardButton) error {
	return editInlineKbd(callback, keyboard)
}

// EditLiveLocation is a method that allows to edit the Location (and others options)
// ONLY for messages sent by the bot that contain it (like Photo or Document...)
func (callback CallbackQuery) EditLiveLocation(latitude, longitude float64, opts *echotron.EditLocationOptions) error {
	return editLiveLocation(callback, latitude, longitude, opts)
}

// EditCaption is a method that allows to edit the caption (and others options)
// ONLY for messages sent by the bot that contain media (like Photo or Document...)
func (callback CallbackQuery) EditCaption(opts *echotron.MessageCaptionOptions) error {
	return editCaption(callback, opts)
}

// Delete  is a method that allows to delete the original message
func (callback CallbackQuery) Delete() error {
	return delete(callback)
}
