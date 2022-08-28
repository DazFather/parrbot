package tgui_test

import (
	"fmt"

	"github.com/DazFather/parrbot/tgui"

	"github.com/NicoNex/echotron/v3"
)

func ExampleParseModeOpt() {
	var opt *tgui.EditOptions = tgui.DisableWebPagePreview(nil)
	fmt.Println(tgui.ParseModeOpt(opt, "HTML").ParseMode) // Output: HTML
	fmt.Println(opt.ParseMode)                            // HTML
	fmt.Println(tgui.ParseModeOpt(nil, "HTML").ParseMode) // HTML
}

func ExampleEntitiesOpt() {
	var (
		fakeEntities        = []echotron.MessageEntity{{URL: "https://go.dev", Language: "Go"}}
		firstEntityLanguage = func(opts *tgui.EditOptions) string {
			return opts.Entities[0].Language
		}
	)

	var opt *tgui.EditOptions = tgui.DisableWebPagePreview(nil)
	fmt.Println(firstEntityLanguage(tgui.EntitiesOpt(opt, fakeEntities))) // Output: Go
	fmt.Println(firstEntityLanguage(opt))                                 // Go
	fmt.Println(firstEntityLanguage(tgui.EntitiesOpt(nil, fakeEntities))) // Go
}

func ExampleReplyMarkupOpt() {
	var markup echotron.InlineKeyboardMarkup = tgui.InlineKeyboard([][]tgui.InlineButton{
		{{Text: "Click me", CallbackData: "/click"}, {Text: "Golang site", URL: "https://go.dev"}},
		{{Text: "Hello World", CallbackData: "/hello"}},
	})

	var opt *tgui.EditOptions = tgui.ParseModeOpt(nil, "HTML")
	fmt.Println(tgui.ReplyMarkupOpt(opt, markup).ReplyMarkup) // Output:
	fmt.Println(opt.ReplyMarkup)                              // HTML
	fmt.Println(tgui.ReplyMarkupOpt(nil, markup).ReplyMarkup) // HTML
}

func ExampleInlineKbdOpt() {
	var (
		keyboard = [][]tgui.InlineButton{
			{{Text: "Click me", CallbackData: "/click"}, {Text: "Golang site", URL: "https://go.dev"}},
			{{Text: "Hello World", CallbackData: "/hello"}},
		}
		firstButtonText = func(opt *tgui.EditOptions) string {
			return opt.ReplyMarkup.InlineKeyboard[0][0].Text
		}
	)

	var opt *tgui.EditOptions = tgui.DisableWebPagePreview(nil)
	fmt.Println(firstButtonText(tgui.InlineKbdOpt(opt, keyboard))) // Output: Click me
	fmt.Println(firstButtonText(opt))                              // Click me
	fmt.Println(firstButtonText(tgui.InlineKbdOpt(nil, keyboard))) // Click me
}

func ExampleDisableWebPagePreview() {
	var opt *tgui.EditOptions = tgui.ParseModeOpt(nil, "HTML")
	fmt.Println(tgui.DisableWebPagePreview(opt).DisableWebPagePreview) // Output: true
	fmt.Println(opt.DisableWebPagePreview)                             // true
	fmt.Println(tgui.DisableWebPagePreview(nil).DisableWebPagePreview) // true
}

func ExampleWebPagePreview() {
	var opt *tgui.EditOptions = tgui.ParseModeOpt(nil, "HTML")
	fmt.Println(tgui.WebPagePreviewOpt(opt, false).DisableWebPagePreview) // Output: true
	fmt.Println(opt.DisableWebPagePreview)                                // true
	fmt.Println(tgui.WebPagePreviewOpt(nil, false).DisableWebPagePreview) // true
}

func ExampleToMessageOptions() {
	var (
		editOpt = &tgui.EditOptions{
			ParseMode:             "HTML",
			DisableWebPagePreview: true,
		}
		opt = tgui.ToMessageOptions(editOpt)
	)

	fmt.Println(tgui.ToMessageOptions(nil))                               // Output: nil
	fmt.Println(editOpt.ParseMode, opt.ParseMode)                         // HTML HTML
	fmt.Println(editOpt.DisableWebPagePreview, opt.DisableWebPagePreview) // true, true
}
