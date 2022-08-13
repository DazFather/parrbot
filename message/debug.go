package message

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NicoNex/echotron/v3"
)

// Log is a useful function to show what values the data is carrying using JSON.
// Tips: Be careful to who you are sending the message or the end user could be
// a bit confused. If you are the developer use your own chatID
func Log(chatID int64, any ...interface{}) {
	var message = Text{
		fmt.Sprint("🦜 <b>Log</b> [", time.Now(), "]\n"),
		&echotron.MessageOptions{ParseMode: echotron.HTML},
	}

	// Parsing each data and add the result to the message text
	for i, value := range any {
		t := fmt.Sprint("\n<b>Data (", i, "):</b>\nString: <code>", strings.ReplaceAll(fmt.Sprint(value), "<nil>", "nil"), "</code>")
		if data, e := json.MarshalIndent(value, "", "   "); e == nil {
			message.Text += fmt.Sprint(t, "\nJSON:\n<code>", string(data), "</code>\n")
		} else {
			message.Text += fmt.Sprint(t, "\n<code>[Impossible to parse JSON]</code>\n")
		}
	}

	// Clipping useful links
	message.ClipInlineKeyboard([][]echotron.InlineKeyboardButton{
		{
			{Text: "📖 Parr(B)ot doc", URL: "https://pkg.go.dev/github.com/DazFather/parrbot"},
			{Text: "🧑‍💻 Parr(B)ot dev", URL: "t.me/DazFather"},
		},
		{
			{Text: "📖 Echotron doc", URL: "https://pkg.go.dev/github.com/NicoNex/echotron/v3"},
			{Text: "👥 Echotron group", URL: "t.me/echotron"},
		},
		{{Text: "📖 Telegram Bot API doc", URL: "https://core.telegram.org/bots/api"}},
	})

	// Send the message to the specified user
	message.Send(chatID)
}
