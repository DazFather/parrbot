package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import (
	"fmt"
	"strconv"
	"strings"

	"parrbot/message"
	"parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

// Menu is a collection of message that can be seen by pressing the related
// button. The menu can be triggered by both a message and a CallbackData.
type Menu struct {
	// The collection of pages of the menu, required to work
	Pages []MenuPage

	// The caption of the navigation buttons. You can embed "[INDEX]" inside
	// PreviousCaption and NextCaption to show the number of the related page.
	// By default (if missing or empty string are passed) their values is:
	// "Next ⏭ [INDEX]", "[INDEX] ⏮ Prev.", "❌ Close".
	NextCaption, PreviousCaption, CloseCaption string
}

// MenuPage is a function that will return the content that will be shown when a user request that page of the Menu
type MenuPage func(b *robot.Bot) (content string, opts *echotron.MessageTextOptions)

// UseMenu implements a Menu to work with a specific trigger (you can direclt add the return to the command list)
func (m *Menu) UseMenu(name, trigger string) robot.Command {
	// Set default captions values
	if m.NextCaption == "" {
		m.NextCaption = "Next ⏭ [INDEX]"
	}
	if m.PreviousCaption == "" {
		m.PreviousCaption = "[INDEX] ⏮ Prev."
	}
	if m.CloseCaption == "" {
		m.CloseCaption = "❌ Close"
	}

	// Create the handler function
	var menuHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
		// Get the payload
		var text string
		if update.Message != nil {
			text = update.Message.Text
			update.Message.Delete()
		} else {
			text = update.CallbackQuery.Data
		}
		text = strings.TrimSpace(strings.TrimPrefix(text, trigger))

		// Select the page of the menu
		switch text {
		case "":
			return m.SelectPage(trigger, 0, bot, update)

		case "x":
			if update.CallbackQuery != nil {
				var api = message.GetAPI()
				api.DeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.ID)
			}

		default:
			n, err := strconv.Atoi(text)
			if n >= 0 && n < len(m.Pages) && err == nil {
				return m.SelectPage(trigger, n, bot, update)
			}
		}

		return nil
	}

	// Create the command and return it
	return robot.Command{
		Description: name,
		Trigger:     trigger,
		ReplyAt:     message.CALLBACK_QUERY + message.MESSAGE,
		CallFunc:    menuHandler,
	}
}

// SelectPage select and call the return of a specific page of the Menu
func (m Menu) SelectPage(trigger string, pageNumber int, b *robot.Bot, u *message.Update) message.Any {
	var content, opt = m.Pages[pageNumber](b)
	if opt == nil {
		opt = new(echotron.MessageTextOptions)
	}
	opt.ReplyMarkup.InlineKeyboard = append(opt.ReplyMarkup.InlineKeyboard, m.genButtons(trigger, pageNumber))

	if u.Message != nil {
		msg := message.Text{
			Text: content,
			Opts: &echotron.MessageOptions{
				ParseMode:             opt.ParseMode,
				Entities:              opt.Entities,
				DisableWebPagePreview: opt.DisableWebPagePreview,
				ReplyMarkup:           opt.ReplyMarkup,
			},
		}
		return msg
	}

	msgID := echotron.NewMessageID(u.CallbackQuery.From.ID, u.CallbackQuery.Message.ID)
	message.GetAPI().EditMessageText(content, msgID, opt)
	return nil

}

// genButtons returns the navigation buttons row (previous / close / next)
func (m Menu) genButtons(trigger string, pageNumber int) (btnRow []echotron.InlineKeyboardButton) {
	// Previous Button
	if pageNumber > 0 {
		btnRow = append(btnRow, echotron.InlineKeyboardButton{
			Text:         strings.ReplaceAll(m.PreviousCaption, "[INDEX]", strconv.Itoa(pageNumber)),
			CallbackData: fmt.Sprint(trigger, " ", pageNumber-1),
		})
	}

	// Close Button
	btnRow = append(btnRow, echotron.InlineKeyboardButton{
		Text:         m.CloseCaption,
		CallbackData: fmt.Sprint(trigger, " x"),
	})

	// Next Button
	if pageNumber < len(m.Pages)-1 {
		btnRow = append(btnRow, echotron.InlineKeyboardButton{
			Text:         strings.ReplaceAll(m.NextCaption, "[INDEX]", strconv.Itoa(pageNumber+2)),
			CallbackData: fmt.Sprint(trigger, " ", pageNumber+1),
		})
	}
	return
}

// StaticPage returns a MenuPage that will return always the same output
func StaticPage(content string, opt *echotron.MessageTextOptions) MenuPage {
	return func(*robot.Bot) (string, *echotron.MessageTextOptions) {
		return content, opt
	}
}
