package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import (
	"errors"
	"log"

	"github.com/DazFather/parrbot/message"
	"github.com/DazFather/parrbot/robot"
)

// ShowMessage allows to show a text message editing the incoming in case of CALLBACK_QUERY or sending a new one otherwhise
func ShowMessage(u message.Update, text string, opt *EditOptions) (sent *message.UpdateMessage, err error) {
	if callback := u.CallbackQuery; callback != nil {
		err = callback.EditText(text, opt)
		if err == nil {
			_, err = callback.Answer(nil)
			sent = callback.Message
		}
		return
	}

	if original := u.FromMessage(); original != nil && original.Chat != nil {
		return message.Text{text, ToMessageOptions(opt)}.Send(original.Chat.ID)
	}

	return nil, errors.New("Invalid given update")
}

// Sender creates a robot.CommandFunc that will always send the given message
func Sender(msg message.Any) robot.CommandFunc {
	return func(bot *robot.Bot, update *message.Update) message.Any {
		return msg
	}
}

// Replier creates a robot.Command that will use the ShowMessage function showing
// a new message or editing the previous in case of message.CALLBACK_QUERY
func Replier(trigger, description string, text string, opt *EditOptions) robot.Command {
	return robot.Command{
		Description: description,
		Trigger:     trigger,
		ReplyAt:     message.MESSAGE + message.CALLBACK_QUERY,
		CallFunc: func(bot *robot.Bot, update *message.Update) message.Any {
			if _, err := ShowMessage(*update, text, opt); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}
