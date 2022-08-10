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
	//CommandList []Command
	ChatID int64 // ChatID of the user who is using the bot on a private chat
}

// Creates a new bot - will be called when a user first start the bot
func newBot(chatID int64) echotron.Bot {
	bot := &Bot{chatID}
	go bot.selfDestruct(time.After(time.Hour * 2))
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

// Start give life to your amazing root parrot
func Start(commandList []Command) {
	// Initialization
	LoadToken()
	message.LoadAPI(TOKEN)
	LoadCommands(commandList)
	dsp = echotron.NewDispatcher(TOKEN, newBot)

	// Put life into the bot
	log.Println(dsp.Poll())
}
