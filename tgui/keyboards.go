package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import "github.com/NicoNex/echotron/v3"

// InlineButton is a shorter way to indicate an echotron.InlineKeyboardButton
type InlineButton = echotron.InlineKeyboardButton

// KeyButton is a shorter way to indicate echotron.KeyboardButton
type KeyButton = echotron.KeyboardButton

// InlineKeyboard allows to quickly cast a matrix of inline buttons into a ReplyMarkup
func InlineKeyboard(kbd [][]InlineButton) echotron.InlineKeyboardMarkup {
	return echotron.InlineKeyboardMarkup{InlineKeyboard: kbd}
}

// Keyboard allows to quickly cast a matrix of buttons into a ReplyMarkup setting some useful options. Arguments:
// disposable - after user tap on any button the keyboard will be removed,
// inputPlaceholder - a text (max 64 chars) that will apper os placeholder on input field
// kbd - the matrix of buttons that compose the keyboard
// Resulting ReplyKeyboardMarkup will also have ResizeKeyboard and Selective set to true
func Keyboard(disposable bool, inputPlaceholder string, kbd [][]KeyButton) echotron.ReplyKeyboardMarkup {
	return echotron.ReplyKeyboardMarkup{
		OneTimeKeyboard:       disposable,
		InputFieldPlaceholder: inputPlaceholder,
		Keyboard:              kbd,
		// Resize the keyboard vertically for optimal fit
		ResizeKeyboard: true,
		// If message is sent on reply or @user is mentioned keyboard will appear just for him
		Selective: true,
	}
}

// KeyboardRemover allows to remove a previouly sent keyboard.
// If globally is false then kwyboard will be removed for user mentioned on the message or on the reply
func KeyboardRemover(globally bool) echotron.ReplyKeyboardRemove {
	return echotron.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      !globally,
	}
}

// Arrange the layout of given buttons in a given number of columns
func Arrange[AnyButton InlineButton | KeyButton](columns int, fromList ...AnyButton) (keyboard [][]AnyButton) {
	if columns < 1 {
		return
	}

	var (
		rowSize int = len(fromList) / columns
		size    int = rowSize * columns
	)

	for i := 0; i < size; i += columns {
		keyboard = append(keyboard, fromList[i:i+columns])
	}

	if len(fromList)-size != 0 {
		keyboard = append(keyboard, fromList[size:])
	}

	return
}

// GenInlineKeyboard allows to generate an echotron.ReplyMarkup from a list of inline buttons
// WARNING: Deprecated
func GenInlineKeyboard(columns int, fromList ...InlineButton) echotron.InlineKeyboardMarkup {
	return InlineKeyboard(Arrange(columns, fromList...))
}
