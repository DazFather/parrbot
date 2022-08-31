package tgui // TeleGram User Interface or Toolkit for Graphical User Interface

import "github.com/NicoNex/echotron/v3"

// EditOptions is a shorter way to indicate an echotron.MessageTextOptions
type EditOptions = echotron.MessageTextOptions

// safeReference return a pointer to the given opt or to a new EditOptions if given is nil
func safeReference(opt *EditOptions) *EditOptions {
	if opt == nil {
		return new(EditOptions)
	}
	return opt
}

// ParseModeOpt allows to change the ParseMode when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func ParseModeOpt(opt *EditOptions, value echotron.ParseMode) *EditOptions {
	opt = safeReference(opt)
	opt.ParseMode = value
	return opt
}

// EntitiesOpt allows to change the Entities when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func EntitiesOpt(opt *EditOptions, value []echotron.MessageEntity) *EditOptions {
	opt = safeReference(opt)
	opt.Entities = value
	return opt
}

// ReplyMarkupOpt allows to change the ReplyMarkup when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func ReplyMarkupOpt(opt *EditOptions, value echotron.InlineKeyboardMarkup) *EditOptions {
	opt = safeReference(opt)
	opt.ReplyMarkup = value
	return opt
}

// InlineKbdOpt allows to change the Entities when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func InlineKbdOpt(opt *EditOptions, value [][]InlineButton) *EditOptions {
	return ReplyMarkupOpt(opt, echotron.InlineKeyboardMarkup{InlineKeyboard: value})
}

// DisableWebPagePreview allows to set DisableWebPagePreview to true when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func DisableWebPagePreview(opt *EditOptions) *EditOptions {
	return WebPagePreviewOpt(opt, false)
}

// WebPagePreviewOpt allows to change the WebPagePreview when editing a message.
// When given the function will edit opt and return it or else it will create a new one
func WebPagePreviewOpt(opt *EditOptions, enable bool) *EditOptions {
	opt = safeReference(opt)
	opt.DisableWebPagePreview = !enable
	return opt
}

// ToMessageOptions transform a given EditOptions into a echotron.MessageOptions
func ToMessageOptions(opt *EditOptions) *echotron.MessageOptions {
	if opt == nil {
		return nil
	}
	return &echotron.MessageOptions{
		ParseMode:             opt.ParseMode,
		Entities:              opt.Entities,
		ReplyMarkup:           opt.ReplyMarkup,
		DisableWebPagePreview: opt.DisableWebPagePreview,
	}
}

// NewAlert generates the option needed to display an alert when answering a
// callback using the Answer method of a CallbackQuery
func NewAlert(text string, cacheTime uint16) *echotron.CallbackQueryOptions {
	return &echotron.CallbackQueryOptions{
		Text:      text,
		CacheTime: int(cacheTime),
		ShowAlert: true,
	}
}

// NewToast generates the option needed to display an toast notification when
// answering a callback using the Answer method of a CallbackQuery
func NewToast(text string, cacheTime uint16) *echotron.CallbackQueryOptions {
	return &echotron.CallbackQueryOptions{
		Text:      text,
		CacheTime: int(cacheTime),
	}
}
