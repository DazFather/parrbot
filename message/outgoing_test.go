package message_test

import (
	"fmt"

	"github.com/DazFather/parrbot/message"

	"github.com/NicoNex/echotron/v3"
)

func ExampleClipKeyboard() {
	var msg = message.Text{"Hello World", nil}

	ptr := msg.ClipKeyboard(echotron.ReplyKeyboardMarkup{
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "Tap on the button below to send",
		Keyboard: [][]echotron.KeyboardButton{{
			{Text: "Useless button"},
		}},
		ResizeKeyboard: true,
	})

	fmt.Println(ptr == &msg) // Output: true
}

func ExampleClipInlineKeyboard() {
	var msg = message.Text{"Hello World", nil}

	ptr := msg.ClipInlineKeyboard([][]echotron.InlineKeyboardButton{{
		{Text: "Go", URL: "https://go.dev"},
	}})

	fmt.Println(ptr == &msg) // Output: true
}
