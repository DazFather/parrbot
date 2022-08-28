package robot

import (
	"log"
	"time"

	"github.com/DazFather/parrbot/message"

	"github.com/NicoNex/echotron/v3"
)

var dsp *echotron.Dispatcher

// Bot structure
type Bot struct {
	ChatID int64 // ChatID of the user who is using the bot on a private chat
}

// Creates a new bot - will be called when a user first start the bot
func newBot(chatID int64) echotron.Bot {
	bot := &Bot{chatID}
	if duration := Config.DeleteSessionTimer; duration != 0 {
		go bot.selfDestruct(time.After(duration))
	}
	return bot
}

func (b *Bot) selfDestruct(timech <-chan time.Time) {
	<-timech
	dsp.DelSession(b.ChatID)
}

// Update is used to manage the incoming inputs from Telegram
func (b *Bot) Update(u *echotron.Update) {
	var update = message.CastUpdate(u)

	fn := Select(update)
	if fn == nil {
		return
	}

	if msg := fn(b, update); msg != nil {
		msg.Send(b.ChatID)
	}
}

// Start give life to your amazing robo-parrot. It accepts the commands that the
// bot will need to handle. This function will also load all configuartion available
// int the Confing variable, so if you want to change them use do it before you
// call this function. Start will also stop the flow of execution (unless bot crash)
// if any error happens it will be logged on screen
// This function will also load configurations and block the flow of execution until the bot crash
func Start(commandList ...Command) {
	// Initialization
	if err := Config.init(); err != nil {
		log.Fatal("Config error: ", err)
	}
	LoadCommands(commandList)
	dsp = echotron.NewDispatcher(Config.token, newBot)

	// Put life into the bot
	log.Println(dsp.Poll())
}
