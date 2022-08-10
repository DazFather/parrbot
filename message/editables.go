package message

import (
	"encoding/json"
	"errors"

	"github.com/NicoNex/echotron/v3"
)

type editable interface {
	extractID() *echotron.MessageIDOptions
	grabMessage() *UpdateMessage
}

type editFn func(echotron.MessageIDOptions) (echotron.APIResponseMessage, error)

func edit(e editable, call editFn) (err error) {
	var msgID = e.extractID()
	if msgID == nil {
		return errors.New("Invalid message or Id")
	}

	// Perform the edit and clearig the response
	var edited *UpdateMessage
	edited, err = clearResponse(call(*msgID))
	if err != nil || edited != nil {
		return
	}

	// Sync changes
	if original := e.grabMessage(); original != nil {
		var jsonData []byte
		jsonData, err = json.Marshal(*edited)
		if err == nil {
			err = json.Unmarshal(jsonData, original)
		}
	}

	return
}

func editText(e editable, text string, opts *echotron.MessageTextOptions) error {
	return edit(e, func(msgID echotron.MessageIDOptions) (echotron.APIResponseMessage, error) {
		return api.EditMessageText(text, *e.extractID(), opts)
	})
}

func editMedia(e editable, media echotron.InputMedia, opts *echotron.MessageReplyMarkup) error {
	return edit(e, func(msgID echotron.MessageIDOptions) (echotron.APIResponseMessage, error) {
		return api.EditMessageMedia(*e.extractID(), media, opts)
	})
}

func editInlineKbd(e editable, keyboard [][]echotron.InlineKeyboardButton) error {
	var opts = &echotron.MessageReplyMarkup{ReplyMarkup: echotron.InlineKeyboardMarkup{InlineKeyboard: keyboard}}

	return edit(e, func(msgID echotron.MessageIDOptions) (echotron.APIResponseMessage, error) {
		return api.EditMessageReplyMarkup(*e.extractID(), opts)
	})
}

func editLiveLocation(e editable, latitude, longitude float64, opts *echotron.EditLocationOptions) error {
	return edit(e, func(msgID echotron.MessageIDOptions) (echotron.APIResponseMessage, error) {
		return api.EditMessageLiveLocation(*e.extractID(), latitude, longitude, opts)
	})
}

func editCaption(e editable, opts *echotron.MessageCaptionOptions) error {
	return edit(e, func(msgID echotron.MessageIDOptions) (echotron.APIResponseMessage, error) {
		return api.EditMessageCaption(*e.extractID(), opts)
	})
}

func delete(e editable) error {
	message := e.grabMessage()
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Unable to retreive message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}

	// Deleting message and clearing response
	res, err := api.DeleteMessage(message.Chat.ID, message.ID)
	if err != nil {
		return ResponseError{"Echotron", 1, err.Error()}
	}
	if !res.Ok {
		return ResponseError{"Telegram", res.ErrorCode, res.Description}
	}
	return nil
}

/* --- Implementing UpdateMessage --- */

func (message *UpdateMessage) grabMessage() *UpdateMessage {
	return message
}

func (message *UpdateMessage) extractID() *echotron.MessageIDOptions {
	if message == nil || message.Chat == nil {
		return nil
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)
	return &msgIDOpt
}

/* --- Implementing CallbackQuery --- */

func (callback CallbackQuery) grabMessage() *UpdateMessage {
	return callback.Message
}

func (callback CallbackQuery) extractID() (msgIDOpt *echotron.MessageIDOptions) {
	msgIDOpt = callback.Message.extractID()
	if msgIDOpt == nil {
		idOpt := echotron.NewInlineMessageID(callback.InlineMessageID)
		msgIDOpt = &idOpt
	}
	return
}

/* --- Implementing Update --- */

func (update Update) grabMessage() *UpdateMessage {
	if update.Message != nil {
		return update.Message
	} else if update.EditedMessage != nil {
		return update.EditedMessage
	} else if update.ChannelPost != nil {
		return update.ChannelPost
	} else if update.EditedChannelPost != nil {
		return update.EditedChannelPost
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.grabMessage()
	}

	return nil
}

func (update Update) extractID() (msgIDOpt *echotron.MessageIDOptions) {
	if callback := update.CallbackQuery; callback != nil {
		msgIDOpt = callback.extractID()
	} else if result := update.ChosenInlineResult; result != nil {
		id := echotron.NewInlineMessageID(result.InlineMessageID)
		msgIDOpt = &id
	} else if msg := update.grabMessage(); msg != nil {
		msgIDOpt = msg.extractID()
	}
	return
}

/* --- Implementing Reference --- */

func (ref Reference) grabMessage() *UpdateMessage {
	return nil
}

func (ref Reference) extractID() (msgIDOpt *echotron.MessageIDOptions) {
	return &ref.MessageIDOptions
}
