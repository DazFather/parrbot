package main

import (
	"parrbot/message"
	"parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

func main() {
	var commandList = []robot.Command{
		{Description: "Start the bot", Trigger: "/start", ReplyAt: message.MESSAGE, CallFunc: helloHandler},
		{Trigger: "/info", ReplyAt: message.CALLBACK_QUERY, CallFunc: infoHandler},
	}

	robot.Start(commandList)
}

var helloHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
	var kbd = [][]echotron.InlineKeyboardButton{{
		{Text: "ℹ️ more info", CallbackData: "/info"},
	}}

	var msg = message.Text{"🦜 Hello World!", nil}
	return *msg.ClipInlineKeyboard(kbd)
}

// this is a valid robot.CommandFunc too (it just needs the right params)
func infoHandler (bot *robot.Bot, update *message.Update) message.Any {
	update.CallbackQuery.EditText("Made with ❤️ by @DazFather", nil)
	update.CallbackQuery.Answer(nil)
	return nil
}
