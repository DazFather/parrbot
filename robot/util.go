package robot

import "github.com/DazFather/parrbot/message"

// Sender creates a CommandFunc that only returns the given message
func Sender(msg message.Any) CommandFunc {
	return func(bot *Bot, update *message.Update) message.Any {
		return msg
	}
}
