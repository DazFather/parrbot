package message

import (
	"github.com/NicoNex/echotron/v3"
)

/* --- Caption --- */

// EditCaption is a method that allows to edit the caption (and others options)
// ONLY for messages sent by the bot that contain media (like Photo or Document...)
func (message *UpdateMessage) EditCaption(opts *echotron.MessageCaptionOptions) (err error) {
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editCaption(msgIDOpt, opts)
	if err == nil {
		message = newMsg
	}

	return
}

func editCaption(msgIDOpt echotron.MessageIDOptions, opts *echotron.MessageCaptionOptions) (message *UpdateMessage, err error) {
	// Perform the edit and clearig the response
	message, err = clearResponse(api.EditMessageCaption(msgIDOpt, opts))

	if err != nil || message != nil {
		return
	}

	// Sync message
	message.Text = opts.Caption
	message.Entities, message.Media.CaptionEntities = nil, nil
	for _, entity := range opts.CaptionEntities {
		message.Entities = append(message.Entities, &entity)
		message.Media.CaptionEntities = append(message.Media.CaptionEntities, &entity)
	}
	message.Media.Caption = opts.Caption
	message.InlineKeyboard = &opts.ReplyMarkup

	return
}

/* --- Live Location --- */

// EditLiveLocation is a method that allows to edit the Location (and others options)
// ONLY for messages sent by the bot that contain it (like Photo or Document...)
func (message *UpdateMessage) EditLiveLocation(latitude, longitude float64, opts *echotron.EditLocationOptions) (err error) {
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editLiveLocation(msgIDOpt, latitude, longitude, opts)
	if err == nil {
		message = newMsg
	}

	return
}

func editLiveLocation(msgIDOpt echotron.MessageIDOptions, latitude, longitude float64, opts *echotron.EditLocationOptions) (message *UpdateMessage, err error) {
	// Perform the edit and clearig the response
	message, err = clearResponse(api.EditMessageLiveLocation(msgIDOpt, latitude, longitude, opts))

	if err != nil || message != nil {
		return
	}

	// Sync message
	message.InlineKeyboard = &opts.ReplyMarkup
	// WARNING: Unable to sync the message (*UpdateMessage) if not returned by EditMessageLiveLocation

	return
}

/* --- Media --- */

// EditMedia is a method that allows to edit the media (and others options)
// ONLY for messages sent by the bot that contain it
func (message *UpdateMessage) EditMedia(media echotron.InputMedia, keyboard [][]echotron.InlineKeyboardButton) (err error) {
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editMedia(msgIDOpt, media, keyboard)
	if err == nil {
		message = newMsg
	}

	return
}

func editMedia(msgIDOpt echotron.MessageIDOptions, media echotron.InputMedia, keyboard [][]echotron.InlineKeyboardButton) (message *UpdateMessage, err error) {
	// Perform the edit and clearig the response
	var opts = &echotron.MessageReplyMarkup{ReplyMarkup: echotron.InlineKeyboardMarkup{InlineKeyboard: keyboard}}
	message, err = clearResponse(api.EditMessageMedia(msgIDOpt, media, opts))

	if err != nil || message != nil {
		return
	}

	// Sync message
	message.InlineKeyboard = &opts.ReplyMarkup
	// WARNING: : Unable to sync the message (*UpdateMessage) with the new media (echotron.InputMedia) if not returned by EditMessageMedia

	return
}

/* --- Inline keyboard --- */

// EditInlineKeyboard is a method that allows to edit the InlineKeyboard ONLY
// for messages sent by the bot
func (message *UpdateMessage) EditInlineKeyboard(keyboard [][]echotron.InlineKeyboardButton) (err error) {
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editInlineKeyboard(msgIDOpt, keyboard)
	if err == nil {
		message = newMsg
	}

	return
}

func editInlineKeyboard(msgIDOpt echotron.MessageIDOptions, keyboard [][]echotron.InlineKeyboardButton) (message *UpdateMessage, err error) {
	// Perform the edit and clearig the response
	var opts = &echotron.MessageReplyMarkup{ReplyMarkup: echotron.InlineKeyboardMarkup{InlineKeyboard: keyboard}}
	message, err = clearResponse(api.EditMessageReplyMarkup(msgIDOpt, opts))

	if err != nil || message != nil {
		return
	}

	// Sync message
	message.InlineKeyboard = &opts.ReplyMarkup

	return
}

/* --- Text --- */

// EditText is a method that allows to edit the text (and others options)
// for textual messages (message.Text) sent by the bot
func (message *UpdateMessage) EditText(text string, opts *echotron.MessageTextOptions) (err error) {
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
	}
	if message.Chat == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid chat ID"}
	}
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	newMsg := new(UpdateMessage)
	newMsg, err = editText(msgIDOpt, text, opts)
	if err == nil {
		message = newMsg
	}

	return
}

func editText(msgIDOpt echotron.MessageIDOptions, text string, opts *echotron.MessageTextOptions) (message *UpdateMessage, err error) {
	// Perform the edit and clearig the response
	message, err = clearResponse(api.EditMessageText(text, msgIDOpt, opts))

	if err != nil || message != nil {
		return
	}

	// Sync message
	message = new(UpdateMessage)
	message.Text = text
	message.Entities = nil
	for _, entity := range opts.Entities {
		message.Entities = append(message.Entities, &entity)
	}
	message.InlineKeyboard = &opts.ReplyMarkup

	return
}

/* --- Delete --- */

// Delete the given message on the original chat and memory (setting it to nil)
func (message *UpdateMessage) Delete() error {
	// Guard closes
	if message == nil {
		return ResponseError{"Parr(B)ot", 1, "Invalid message"}
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
	message = nil

	return nil
}
