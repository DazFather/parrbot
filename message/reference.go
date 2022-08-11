package message

import (
	"github.com/NicoNex/echotron/v3"
)

// Reference rapresent a reference to an existing message. You can use it to do
// the same actions (editing or deleting) but with a lighter structure that can
// used as a echotron.MessageIDOptions thanks to embedding
type Reference struct {
	messageID int
	chatID    int64
	echotron.MessageIDOptions
}

// NewReference creates a new Reference that is referencing the message contained inside e
func NewReference(e editable) *Reference {
	var (
		msgID   = e.extractID()
		message = e.grabMessage()
	)
	if msgID == nil || message == nil || message.Chat == nil {
		return nil
	}

	return &Reference{
		messageID:        message.ID,
		chatID:           message.Chat.ID,
		MessageIDOptions: *msgID,
	}
}

// MessageID allows to get the messageID of the referenced message
func (ref Reference) MessageID() int {
	return ref.messageID
}

// ChatID allows to get the ChatID in witch referenced message was sent
func (ref Reference) ChatID() int64 {
	return ref.chatID
}

// EditText is a method that allows to edit the text (and others options)
// for textual messages (message.Text)
func (ref Reference) EditText(text string, opts *echotron.MessageTextOptions) error {
	return editText(ref, text, opts)
}

// EditMedia is a method that allows to edit the media (and others options)
// ONLY for referenced messages that contain it
func (ref Reference) EditMedia(media echotron.InputMedia, opts *echotron.MessageReplyMarkup) error {
	return editMedia(ref, media, opts)
}

// EditInlineKeyboard is a method that allows to edit the InlineKeyboard of
// referenced messages
func (ref Reference) EditInlineKeyboard(keyboard [][]echotron.InlineKeyboardButton) error {
	return editInlineKbd(ref, keyboard)
}

// EditLiveLocation is a method that allows to edit the Location (and others options)
// ONLY for referenced messages that contain it (like Photo or Document...)
func (ref Reference) EditLiveLocation(latitude, longitude float64, opts *echotron.EditLocationOptions) error {
	return editLiveLocation(ref, latitude, longitude, opts)
}

// EditCaption is a method that allows to edit the caption (and others options)
// ONLY for referenced messages that contain media (like Photo or Document...)
func (ref Reference) EditCaption(opts *echotron.MessageCaptionOptions) error {
	return editCaption(ref, opts)
}

func (ref Reference) Delete() error {
	_, err := api.DeleteMessage(ref.chatID, ref.messageID)
	return err
}
