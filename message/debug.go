package message

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NicoNex/echotron/v3"
)

// Log is a useful function to show what values the data is carrying using JSON.
// tips: Be careful to who you are sending the message or the end user could
//       be a bit confused. If you are the developer use your own chatID
func Log(api echotron.API, chatID int64, any ...interface{}) {
	var message = Text{
		fmt.Sprint("🦜 <b>Log</b> [", time.Now(), "]\n"),
		&echotron.MessageOptions{ParseMode: echotron.HTML},
	}

	// Parsing each data and add the result to the message text
	for i, value := range any {
		if data, e := json.MarshalIndent(value, "", "   "); e == nil {
			message.Text = string(data)
		} else {
			message.Text = "Impossible to parse data"
		}
		message.Text = fmt.Sprint("\nData (", i, "):\n<code>", message.Text, "</code>")
	}

	// Clipping useful links
	message.ClipInlineKeyboard([][]echotron.InlineKeyboardButton{{
		{Text: "📖 Echotron", URL: "https://pkg.go.dev/github.com/NicoNex/echotron/v3"},
		{Text: "📖 Telegram", URL: "https://core.telegram.org/bots/api"},
	}})

	// Send the message to the specified user
	message.Send(api, chatID)
}
