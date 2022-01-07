package message

import (
	"github.com/NicoNex/echotron/v3"
)

// CallbackQuery is the Parr(b)ot rapresentation the echotron.CallbackQuery,
// the difference is that it have the custom UpdateMessage instead of echotron.Message type
type CallbackQuery struct {
	ID              string         `json:"id"`
	From            *echotron.User `json:"from"`
	Message         *UpdateMessage `json:"parrbot_message,omitempty"`
	InlineMessageID string         `json:"inline_message_id,omitempty"`
	ChatInstance    string         `json:"chat_instance,omitempty"`
	Data            string         `json:"data,omitempty"`
	GameShortName   string         `json:"game_short_name,omitempty"`
}

func (callback *CallbackQuery) Answer(opts *echotron.CallbackQueryOptions) (resonse bool, err error) {
	res, e := api.AnswerCallbackQuery(callback.ID, opts)
	resonse = res.Result
	switch {
	case e != nil:
		err = ResponseError{"Echotron", 1, err.Error()}
	case !res.Ok:
		err = ResponseError{"Telegram", res.ErrorCode, res.Description}
	}

	return
}

func (callback *CallbackQuery) EditCaption(opts *echotron.MessageCaptionOptions) (err error) {
	// Try to edit the message
	err = callback.Message.EditCaption(opts)
	if err == nil {
		return
	}

	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewInlineMessageID(callback.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editCaption(msgIDOpt, opts)
	if err == nil {
		callback.Message = newMsg
	}

	return
}

func (callback *CallbackQuery) EditLiveLocation(latitude, longitude float64, opts *echotron.EditLocationOptions) (err error) {
	// Try to edit the message
	err = callback.Message.EditLiveLocation(latitude, longitude, opts)
	if err == nil {
		return
	}

	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewInlineMessageID(callback.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editLiveLocation(msgIDOpt, latitude, longitude, opts)
	if err == nil {
		callback.Message = newMsg
	}

	return
}

func (callback *CallbackQuery) EditMedia(media echotron.InputMedia, keyboard [][]echotron.InlineKeyboardButton) (err error) {
	// Try to edit the message
	err = callback.Message.EditMedia(media, keyboard)
	if err == nil {
		return
	}

	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewInlineMessageID(callback.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editMedia(msgIDOpt, media, keyboard)
	if err == nil {
		callback.Message = newMsg
	}

	return
}

func (callback *CallbackQuery) EditInlineKeyboard(keyboard [][]echotron.InlineKeyboardButton) (err error) {
	// Try to edit the message
	err = callback.Message.EditInlineKeyboard(keyboard)
	if err == nil {
		return
	}

	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewInlineMessageID(callback.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editInlineKeyboard(msgIDOpt, keyboard)
	if err == nil {
		callback.Message = newMsg
	}

	return
}

func (callback *CallbackQuery) EditText(text string, opts *echotron.MessageTextOptions) (err error) {
	// Try to edit the message
	err = callback.Message.EditText(text, opts)
	if err == nil {
		return
	}

	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewInlineMessageID(callback.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editText(msgIDOpt, text, opts)
	if err == nil {
		callback.Message = newMsg
	}

	return
}
