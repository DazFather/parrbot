package main

import (
	"Parrbot/message"
	"Parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

func main() {
	var commandList = []robot.Command{
		{Name: "Start", Trigger: "/start", ReplyAt: message.MESSAGE, Scope: helloHandler},
		{Name: "Credits", Trigger: "/info", ReplyAt: message.CALLBACK_QUERY, Scope: infoHandler},
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

var infoHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
	msgID := echotron.NewMessageID(update.CallbackQuery.From.ID, update.CallbackQuery.Message.ID)
	bot.EditMessageText("Made with ❤️ by @DazFather", msgID, nil)
	return nil
}
