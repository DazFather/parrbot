// THE FOLLOWING IS AN EXAMPLE OF A WORKING BOT.
// API TOKEN can be passed as program's argument directly: `<executable> <TOKEN>`,
// or using a file that contains the token: `<executable> --READFROM <filepath>` (case insensitive)
// Another way is to pass it as argument to the function: robot.Config.SetAPIToken
package main

import (
	"github.com/DazFather/parrbot/message" // (Core) Incoming / Outgoing message-related
	"github.com/DazFather/parrbot/robot"   // (Core) Parr(B)ot core functionality
	"github.com/DazFather/parrbot/tgui"    // (Utility) Toolkit for UI
)

func main() {
	// Define the list of commands of the bot
	var commandList = []robot.Command{
		{
			Description: "Start the bot",                          // Text that will appear on bot's menu if ReplyAt contains message.MESSAGE
			Trigger:     "/start",                                 // Trigger that will execute the CallFunc (the first word, MUST start with '/')
			ReplyAt:     message.MESSAGE + message.CALLBACK_QUERY, // What type of update CallFunc is available (you can combine more types using +)
			CallFunc:    startHandler,                             // The function that is going to be executed
		},
		{Trigger: "/info", ReplyAt: message.CALLBACK_QUERY, CallFunc: infoHandler},
		// UseMenu method will generate the needed command for the using the menu
		tgui.UseMenu(helpHandler, "/help", "Help menu"),
	}
	// Make the bot alive
	robot.Start(commandList...)
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
	var callback = update.CallbackQuery
	callback.EditText("Made with ❤️ by @DazFather", nil)
	callback.AnswerToast("Thanks for using me ❤️", 3600)
	return nil // You are NOT obligated to sent a message every time
}

// this is a Menu, is composed by varius pages that you can navigate using previous, close and next buttons
var helpHandler = &tgui.PagedMenu{
	Pages: []tgui.Page{
		// Page 1 - StaticPage is a page that show always the same values
		tgui.StaticPage("No one is gonna help you", nil),
		// Page 2 - Normal pages allows you to interact with the bot for real - time results
		func(b *robot.Bot, u *message.Update) (string, *tgui.EditOptions) {
			return "Just kidding " + u.FromMessage().Chat.FirstName, nil
		},
		// Page 3 - You can always attach options like a custom keyboard
		tgui.StaticPage(
			"Feel free to contact me on Telegram at @DazFather ❤️",
			tgui.InlineKbdOpt(nil, tgui.Arrange(2,
				tgui.InlineLink("Contact developer", "t.me/DazFather"),
				tgui.InlineLink("Echotron group", "t.me/echotron"),
			)),
		),
	},
	// You can also overrite default caption for buttons using CloseCaption, PreviousCaption and...
	NextCaption: "Next page: [INDEX]", // You can use [INDEX] only for next and previous page indexes
}
