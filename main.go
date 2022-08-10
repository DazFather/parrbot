// THE FOLLOWING IS AN EXAMPLE OF A WORKING BOT
// API TOKEN can be passed as program's argument or using --READFROM <filepath> (case insensitive)
package main

import (
	"parrbot/message" // (Core) Incoming / Outgoing message-related
	"parrbot/robot"   // (Core) Parr(B)ot core functionality
	"parrbot/tgui"    // (Utility) Toolkit for UI
)

func main() {
	// Define the list of commands of the bot
	var commandList = []robot.Command{
		{
			Description: "Start the bot",                          // Text that will appear on bot's menu if ReplyAt contains message.MESSAGE
			Trigger:     "/start",                                 // Trigger that will execute the CallFunc (first word)
			ReplyAt:     message.MESSAGE + message.CALLBACK_QUERY, // Wwhat type of update CallFunc is avaiable (you can combine more types using +)
			CallFunc:    startHandler,                             // The function that is going to be executed
		},
		{Trigger: "/info", ReplyAt: message.CALLBACK_QUERY, CallFunc: infoHandler},
		// UseMenu method will generate the needed command for the using the menu
		helpHandler.UseMenu("Help menu", "/help"),
	}
	// Make the bot alive
	robot.Start(commandList)
}

// every robot.CommandFunc can return a new message to be sent, to better organize your code
var startHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
	keyboard := [][]tgui.InlineButton{
		{{Text: "ℹ️ More info", CallbackData: "/info"}, {Text: "🆘 Help", CallbackData: "/help"}},
		{{Text: "Parr(B)ot channel", URL: "t.me/+3_LBajtkqUgzOTFk"}},
	}

	var msg = message.Text{"🦜 Hello World!", nil}
	msg.ClipInlineKeyboard(keyboard)
	return msg
}

// this is another valid robot.CommandFunc too (it just needs the right params)
func infoHandler(bot *robot.Bot, update *message.Update) message.Any {
	update.CallbackQuery.EditText("Made with ❤️ by @DazFather", nil)
	update.CallbackQuery.Answer(nil)
	return nil // You are not obligated to sent a message every time
}

// this is a Menu, is composed by varius pages that you can navigate using previous, close and next buttons
var helpHandler = tgui.Menu{
	Pages: []tgui.MenuPage{
		// Page 1 - StaticPage is a page that show always the same values
		tgui.StaticPage("No one is gonna help you", nil),
		// Page 2 - Normal pages allows you to interact with the bot for real - time results
		func(b *robot.Bot) (string, *tgui.EditOptions) {
			res, _ := message.GetAPI().GetChat(b.ChatID)
			return "Just kidding" + res.Result.FirstName, nil
		},
		// Page 3 - You can always attach options like a custom keyboard
		tgui.StaticPage(
			"Feel free to contact me on Telegram at @DazFather ❤️",
			tgui.InlineKbdOpt(nil, tgui.Arrange(2, []tgui.InlineButton{
				{Text: "Contact developer", URL: "t.me/DazFather"},
				{Text: "Echotron group", URL: "t.me/echotron"},
			}...)),
		),
	},
	// You can also overrite default caption for buttons using CloseCaption, PreviousCaption and...
	NextCaption: "Next page: [INDEX]", // You can use [INDEX] only for next and previous page indexes
}
