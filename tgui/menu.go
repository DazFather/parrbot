package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/DazFather/parrbot/message"
	"github.com/DazFather/parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

// ShowMessage allows to show a text message editing the incoming in case of CALLBACK_QUERY or sending a new one
func ShowMessage(u message.Update, text string, opt *EditOptions) (sent *message.UpdateMessage, err error) {
	if callback := u.CallbackQuery; callback != nil {
		err = callback.EditText(text, opt)
		if err == nil {
			_, err = callback.Answer(nil)
			sent = callback.Message
		}
		return
	}

	if original := u.FromMessage(); original != nil && original.Chat != nil {
		return message.Text{text, ToMessageOptions(opt)}.Send(original.Chat.ID)
	}

	return nil, errors.New("Invalid given update")
}

/* --- Menu --- */

// Menu represents a generic graphical menu introduced by this package (ex. PagedMenu, InlineMenu)
// All menus can be triggered by both a MESSAGE and a CALLBACK_QUERY.
type Menu interface {
	// Initialize the menu before the robot.CommandFunc is created
	initialize(trigger string)

	// Select the previous page, nil if is not possible. Called when payload = "back"
	SelectPrevious() *Page

	// Select the home page of the menu. Called when payload = ""
	SelectHome() Page

	// Select a generic page specified by pageSelector, the payload after trigger.
	// The error will cause an alert (if possible) and a log on console
	Select(pageSelector string) (*Page, error)

	// Show the selected page on screen. If returned the error will be fatal
	Show(page Page, b *robot.Bot, u *message.Update) error
}

// Page is a function that will return the content that will be shown when a user request that page of the Menu
type Page func(bot *robot.Bot, update *message.Update) (content string, opts *EditOptions)

// UseMenu allows to generate the robot.Command from a given menu to make it work
func UseMenu(menu Menu, trigger, description string) robot.Command {
	// Initialize the menu
	menu.initialize(trigger)

	// Create the command handler that will be called at every new update
	var menuHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
		var (
			payload string
			page    Page
		)

		// Extract command payload
		if update.Message != nil {
			payload = update.Message.Text
			update.Message.Delete()
		} else {
			payload = update.CallbackQuery.Data
		}
		payload = strings.TrimSpace(strings.TrimPrefix(payload, trigger))

		// Select menu's page
		switch payload {
		case "":
			page = menu.SelectHome()

		case "x":
			if callback := update.CallbackQuery; callback != nil {
				callback.Delete()
			}
			return nil

		case "back":
			p := menu.SelectPrevious()
			if p == nil {
				collapse(update, "⚠️ Menu expired: please send "+trigger+" again")
				return nil
			}
			page = *p

		default:
			p, err := menu.Select(payload)
			if err != nil {
				collapse(update, "⚠️ Invalid page: retry to send "+trigger)
				log.Println(err)
				return nil
			}
			page = *p
		}

		if err := menu.Show(page, bot, update); err != nil {
			log.Fatal("Error in Show:", err)
		}
		return nil
	}

	// Create the command and return it
	return robot.Command{
		Description: description,
		Trigger:     trigger,
		ReplyAt:     message.CALLBACK_QUERY + message.MESSAGE,
		CallFunc:    menuHandler,
	}
}

// collapse delete the message contained in the update and display an alert with
// given message (if != "") in case of a CALLBACK_QUERY
func collapse(update *message.Update, message string) {
	if callback := update.CallbackQuery; callback != nil {
		var opt *echotron.CallbackQueryOptions
		if message != "" {
			opt = &echotron.CallbackQueryOptions{Text: message, CacheTime: 3600}
		}
		callback.Answer(opt)
		callback.Delete()
	} else {
		update.DeleteMessage()
	}
}

// StaticPage returns a Page that will return always the same output
func StaticPage(content string, pageOption *EditOptions) Page {
	return func(*robot.Bot, *message.Update) (string, *EditOptions) {
		return content, pageOption
	}
}

/* --- Paged Menu --- */

// PagedMenu is a collection of messages in a book-like structure. Every message
// can be seen by pressing the next or previous button
type PagedMenu struct {
	// The collection of pages of the menu, required to work
	Pages []Page

	// The caption of the navigation buttons. You can embed "[INDEX]" inside
	// PreviousCaption and NextCaption to show the number of the related page.
	// By default (if missing or empty string are passed) their values is:
	// "⏭ [INDEX]", "[INDEX] ⏮", "❌".
	NextCaption, PreviousCaption, CloseCaption string
	trigger                                    string
	current                                    int
}

// initialize a PagedMenu setting optional captions, given trigger and current page
func (m *PagedMenu) initialize(trigger string) {
	if m.NextCaption == "" {
		m.NextCaption = "⏭ [INDEX]"
	}
	if m.PreviousCaption == "" {
		m.PreviousCaption = "[INDEX] ⏮"
	}
	if m.CloseCaption == "" {
		m.CloseCaption = "❌"
	}
	m.trigger = trigger
	m.current = 0
}

// SelectPrevious allows to get the page before the current and reset current.
// When not possible (current page index = 0) returns nil
func (m *PagedMenu) SelectPrevious() *Page {
	if m.current <= 0 {
		return nil
	}
	m.current--
	return &m.Pages[m.current]
}

// SelectHome allows to get the first page and reset current
func (m *PagedMenu) SelectHome() Page {
	m.current = 0
	return m.Pages[0]
}

// Select a page by it's index (converting it into an integer) and reset current
func (m *PagedMenu) Select(pageIndex string) (*Page, error) {
	n, err := strconv.Atoi(pageIndex)
	if n >= 0 && n < len(m.Pages) && err != nil {
		return nil, err
	}
	m.current = n
	return &m.Pages[n], nil
}

// Show a given page adding the navigation buttons: previous / close / next
func (m *PagedMenu) Show(page Page, b *robot.Bot, u *message.Update) error {
	var (
		keyboard     [][]InlineButton
		content, opt = page(b, u)
		pageOpt      *EditOptions
	)

	if opt != nil {
		pageOpt = new(EditOptions)
		*pageOpt = *opt
		keyboard = opt.ReplyMarkup.InlineKeyboard
	}

	if size := len(keyboard); size == 0 || !strings.HasPrefix(keyboard[size-1][0].CallbackData, m.trigger+" ") {
		keyboard = append(keyboard, m.genButtons())
	}

	_, err := ShowMessage(*u, content, InlineKbdOpt(pageOpt, keyboard))
	return err
}

// genButtons returns the navigation buttons row (previous / close / next)
func (m PagedMenu) genButtons() (btnRow []InlineButton) {
	// Previous Button
	if m.current > 0 {
		btnRow = append(btnRow, InlineCaller(
			strings.ReplaceAll(m.PreviousCaption, "[INDEX]", strconv.Itoa(m.current)),
			m.trigger,
			strconv.Itoa(m.current-1),
		))
	}

	// Close Button
	btnRow = append(btnRow, InlineCaller(m.CloseCaption, m.trigger, "x"))

	// Next Button
	if m.current < len(m.Pages)-1 {
		btnRow = append(btnRow, InlineCaller(
			strings.ReplaceAll(m.NextCaption, "[INDEX]", strconv.Itoa(m.current+2)),
			m.trigger,
			strconv.Itoa(m.current+1),
		))
	}
	return
}

/* --- Inline Menu --- */

// InlineMenu is a collection of messages in a tree-like structure. Every message
// can be seen by pressing the related button.
type InlineMenu struct {
	// The main page of the menu required to work. All the other pages are nested inside
	Home InlineMenuItem

	// The caption of the inline button that allows user to go back to the previous page
	BackCaption string

	current *InlineMenuItem
	trigger string
	open    *message.Reference
}

// InlineMenuItem rapresent an element of the InlineMenu
type InlineMenuItem struct {
	// The caption of the inline button that when pressed will show Page
	Caption string

	// The actual page that will be shown
	Page Page

	// The pages that will be seen as inline buttons on the bottom when Page will be shown
	Children [][]InlineMenuItem

	parent *InlineMenuItem
}

// initialize an InlineMenu setting BackCaption if not present, given trigger and current page
func (m *InlineMenu) initialize(trigger string) {
	if m.BackCaption == "" {
		m.BackCaption = "🔙 Go back"
	}
	m.trigger = trigger
	m.current = &m.Home
}

// SelectPrevious allows to get the parent page of the current and reset current.
func (m *InlineMenu) SelectPrevious() *Page {
	if m.current == nil || m.current.parent == nil {
		return nil
	}
	current := m.current.parent
	m.current = current
	return &current.Page
}

// SelectHome allows to get the Home's Page and reset current
func (m *InlineMenu) SelectHome() Page {
	var home = &m.Home
	m.current = home
	return home.Page
}

// Select a page by it's indexes (row and column, converting them into integers) and reset current
func (m *InlineMenu) Select(payload string) (*Page, error) {
	if m.current == nil {
		return nil, errors.New("Current page is not setted")
	}

	var (
		values   []string = strings.Split(payload, " ")
		row, col int
		err      error
		selected *InlineMenuItem
	)

	if len(values) != 2 {
		return nil, errors.New("Indexes on payload must be 2")
	}

	row, err = strconv.Atoi(values[0])
	if err != nil {
		return nil, err
	}

	col, err = strconv.Atoi(values[1])
	if err != nil {
		return nil, err
	}

	if len(m.current.Children) <= row || len(m.current.Children[row]) <= col {
		return nil, errors.New("Invalid indexes passed")
	}

	selected = &m.current.Children[row][col]
	if selected == nil {
		return nil, errors.New("Invalid page selected")
	}

	selected.parent = m.current
	m.current = selected

	return &selected.Page, nil
}

// Show a given page adding the buttons for it's Children
func (m *InlineMenu) Show(page Page, b *robot.Bot, u *message.Update) error {
	content, opt := page(b, u)
	opt = InlineKbdOpt(opt, m.current.genKeyboard(m.trigger, m.BackCaption))

	sent, err := ShowMessage(*u, content, opt)
	if err != nil {
		return err
	}

	if m.open == nil {
		m.open = message.NewReference(sent)
	} else if sent.ID != m.open.MessageID() {
		err = m.open.Delete()
		if err == nil {
			m.open = message.NewReference(sent)
		}
	}
	return nil
}

// genKeyboard generate the buttons for the selected item Children
func (i InlineMenuItem) genKeyboard(trigger, backCaption string) (keyboard [][]InlineButton) {
	keyboard = make([][]InlineButton, len(i.Children))

	for i, menuRow := range i.Children {
		row := make([]InlineButton, len(menuRow))
		for j, item := range menuRow {
			row[j] = InlineCaller(item.Caption, trigger, strconv.Itoa(i), strconv.Itoa(j))
		}
		keyboard[i] = row
	}

	if i.parent != nil {
		keyboard = append(keyboard, Wrap(InlineCaller(backCaption, trigger, "back")))
	}

	return
}
