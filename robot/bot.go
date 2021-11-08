package robot

import (
	"log"

	"parrbot/message"

	"github.com/NicoNex/echotron/v3"
)

// Bot structure
type Bot struct {
	ChatID       int64 // ChatID of the user who is using the bot on a private chat
	echotron.API       // Implements all echotron.API methods
}

// Creates a new bot - will be called when a user first start the bot
func newBot(chatID int64) echotron.Bot {
	return &Bot{chatID, echotron.NewAPI(TOKEN)}
}

// Update is used to manage the incoming inputs from Telegram
func (b *Bot) Update(u *echotron.Update) {
	var update = message.CastUpdate(u)

	fn := Select(update)
	if fn == nil {
		return
	}

	if msg := fn(b, update); msg != nil {
		msg.Send(b.API, b.ChatID)
	}
}

// Start give life to your amazing root parrot
func Start(commandList []Command) {
	// Initialization
	LoadToken()
	LoadCommands(commandList)

	// Put life into the bot
	dsp := echotron.NewDispatcher(TOKEN, newBot)
	log.Println(dsp.Poll())
}
