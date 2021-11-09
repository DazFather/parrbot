package message

import (
	"github.com/NicoNex/echotron/v3"
)

// EditCaption is a method that allows to edit the caption (and others options)
// ONLY for messages sent by the bot that contain media (like Photo or Document...)
func (message *UpdateMessage) EditCaption(a echotron.API, opts *echotron.MessageCaptionOptions) (err error) {
	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	// Perform the edit and clearig the response
	var newMsg *UpdateMessage
	newMsg, err = clearResponse(a.EditMessageCaption(msgIDOpt, opts))
	if err != nil {
		return
	}

	// Sync message
	if newMsg != nil {
		message = newMsg
		return
	}
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

// EditLiveLocation is a method that allows to edit the Location (and others options)
// ONLY for messages sent by the bot that contain it (like Photo or Document...)
func (message *UpdateMessage) EditLiveLocation(a echotron.API, latitude, longitude float64, opts *echotron.EditLocationOptions) (err error) {
	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	// Perform the edit and clearig the response
	var newMsg *UpdateMessage
	newMsg, err = clearResponse(a.EditMessageLiveLocation(msgIDOpt, latitude, longitude, opts))
	if err != nil {
		return
	}

	// Sync message
	if newMsg != nil {
		message = newMsg
		return
	}
	message.InlineKeyboard = &opts.ReplyMarkup
	// WARNING: Unable to sync the message (*UpdateMessage) if not returned by EditMessageLiveLocation

	return
}

// EditMedia is a method that allows to edit the media (and others options)
// ONLY for messages sent by the bot that contain it
func (message *UpdateMessage) EditMedia(a echotron.API, media echotron.InputMedia, keyboard [][]echotron.InlineKeyboardButton) (err error) {
	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	// Perform the edit and clearig the response
	var (
		newMsg *UpdateMessage
		opts   = &echotron.MessageReplyMarkup{ReplyMarkup: echotron.InlineKeyboardMarkup{InlineKeyboard: keyboard}}
	)
	newMsg, err = clearResponse(a.EditMessageMedia(msgIDOpt, media, opts))
	if err != nil {
		return
	}

	// Sync message
	if newMsg != nil {
		message = newMsg
		return
	}
	message.InlineKeyboard = &opts.ReplyMarkup
	// WARNING: : Unable to sync the message (*UpdateMessage) with the new media (echotron.InputMedia) if not returned by EditMessageMedia

	return
}

// EditInlineKeyboard is a method that allows to edit the InlineKeyboard ONLY
// for messages sent by the bot
func (message *UpdateMessage) EditInlineKeyboard(a echotron.API, keyboard [][]echotron.InlineKeyboardButton) (err error) {
	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	// Perform the edit and clearig the response
	var (
		newMsg *UpdateMessage
		opts   = &echotron.MessageReplyMarkup{ReplyMarkup: echotron.InlineKeyboardMarkup{InlineKeyboard: keyboard}}
	)
	newMsg, err = clearResponse(a.EditMessageReplyMarkup(msgIDOpt, opts))
	if err != nil {
		return
	}

	// Sync message
	if newMsg != nil {
		message = newMsg
		return
	}
	message.InlineKeyboard = &opts.ReplyMarkup

	return
}

// EditText is a method that allows to edit the text (and others options)
// for textual messages (message.Text) sent by the bot
func (message *UpdateMessage) EditText(a echotron.API, text string, opts *echotron.MessageTextOptions) (err error) {
	// Extracting the MessageIDOptions
	msgIDOpt := echotron.NewMessageID(message.Chat.ID, message.ID)

	// Perform the edit and clearig the response
	var newMsg *UpdateMessage
	newMsg, err = clearResponse(a.EditMessageText(text, msgIDOpt, opts))
	if err != nil {
		return
	}

	// Sync message
	if newMsg != nil {
		message = newMsg
		return
	}
	message.Text = text
	message.Entities = nil
	for _, entity := range opts.Entities {
		message.Entities = append(message.Entities, &entity)
	}
	message.InlineKeyboard = &opts.ReplyMarkup

	return
}
