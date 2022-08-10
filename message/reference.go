package message

import (
	"github.com/NicoNex/echotron/v3"
)

type Reference struct {
	messageID int
	chatID    int64
	echotron.MessageIDOptions
}

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

func (ref Reference) MessageID() int {
	return ref.messageID
}

func (ref Reference) ChatID() int64 {
	return ref.chatID
}

func (ref Reference) Delete() error {
	_, err := api.DeleteMessage(ref.chatID, ref.messageID)
	return err
}
