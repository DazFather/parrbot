package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import (
	"errors"
	"log"
	"regexp"

	"github.com/DazFather/parrbot/message"
	"github.com/DazFather/parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

// ShowMessage allows to show a text message editing the incoming in case of CALLBACK_QUERY or sending a new one otherwhise
func ShowMessage(u message.Update, text string, opt *EditOptions) (sent *message.UpdateMessage, err error) {
	if callback := u.CallbackQuery; callback != nil {
		err = callback.EditText(text, opt)
		if err == nil {
			err = callback.Answer(nil)
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

// Closer creates a robot.Command that will reply only at message.CALLBACK_QUERY
// and given trigger. When it does get called, it deletes the original message
// and shows a toast notification or alert containg the payload of the callback:
// the CallbackQuery.Data without the initial trigger
func Closer(trigger string, alert bool) robot.Command {
	var pattern = regexp.MustCompile(regexp.QuoteMeta(trigger) + `\s*`)

	return robot.Command{
		Trigger: trigger,
		ReplyAt: message.CALLBACK_QUERY,
		CallFunc: func(bot *robot.Bot, update *message.Update) message.Any {
			var callback = update.CallbackQuery

			if text := pattern.ReplaceAllString(callback.Data, ""); text != "" {
				callback.Answer(&echotron.CallbackQueryOptions{
					Text:      text,
					CacheTime: 3600,
					ShowAlert: alert,
				})
			}

			callback.Delete()
			return nil
		},
	}
}
