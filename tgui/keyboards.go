package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import "github.com/NicoNex/echotron/v3"

// InlineButton is a shorter way to indicate an echotron.InlineKeyboardButton
type InlineButton = echotron.InlineKeyboardButton

// InlineKeyboard allows to quickly cast a matrix of inline buttons into a ReplyMarkup
func InlineKeyboard(kbd [][]InlineButton) echotron.InlineKeyboardMarkup {
	return echotron.InlineKeyboardMarkup{InlineKeyboard: kbd}
}

// GenInlineKeyboard allows to generate an echotron.ReplyMarkup from a list of inline buttons
func GenInlineKeyboard(columns int, fromList ...InlineButton) echotron.InlineKeyboardMarkup {
	if columns < 1 {
		return echotron.InlineKeyboardMarkup{}
	}

	var (
		row      []echotron.InlineKeyboardButton
		finalKbd [][]echotron.InlineKeyboardButton
	)

	for _, btn := range fromList {
		row = append(row, btn)
		if len(row) >= columns {
			finalKbd = append(finalKbd, row)
			row = nil
		}
	}
	if len(row) > 0 {
		finalKbd = append(finalKbd, row)
	}

	return InlineKeyboard(finalKbd)
}
