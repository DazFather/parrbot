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

// Start give life to your amazing robo-parrot
func Start(commandList []Command) {
	// Initialization
	if TOKEN == "" {
		token, err := retrieveToken()
		if err != nil {
			log.Fatal(err)
		}
		TOKEN = token
	}
	if err := validateToken(TOKEN); err != nil {
		log.Fatal(err)
	}
	message.LoadAPI(TOKEN)
	LoadCommands(commandList)
	dsp = echotron.NewDispatcher(TOKEN, newBot)

	// Put life into the bot
	log.Println(dsp.Poll())
}
