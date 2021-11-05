package robot

import (
	"log"

	"Parrbot/message"
	"github.com/NicoNex/echotron/v3"
)

// Bot structure
type Bot struct {
	ChatID int64
	echotron.API
}

// Create a new bot
func newBot(chatID int64) echotron.Bot {
	return &Bot{chatID, echotron.NewAPI(TOKEN)}
}

// Manage the incoming inputs (uptate) from Telegram
func (b *Bot) Update(u *echotron.Update) {
	var update = message.Update(*u)

	fn := Select(&update)
	if fn == nil {
		return
	}

	if msg := fn(b, &update); msg != nil {
		msg.Send(b.API, b.ChatID)
	}
}

func Start(commandList []Command) {
	LoadToken()
	LoadCommands(commandList)

	dsp := echotron.NewDispatcher(TOKEN, newBot)
	log.Println(dsp.Poll())
}
