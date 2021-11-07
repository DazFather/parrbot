package robot

import (
	"regexp"
	"strings"

	"Parrbot/message"

	"github.com/NicoNex/echotron/v3"
)

// extractText returns the text contained in the given update.
func extractText(update *echotron.Update) string {
	switch true {
	case update.Message != nil:
		if update.Message.Text != "" {
			return update.Message.Text
		}
		return update.Message.Caption
	case update.EditedMessage != nil:
		return update.EditedMessage.Text
	case update.ChannelPost != nil:
		return update.ChannelPost.Text
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost.Text
	case update.CallbackQuery != nil:
		return update.CallbackQuery.Data
	}

	return ""
}

// extractChatID get the Message.ID of an update (that is not inline-based)
func extractChatID(update *echotron.Update) int64 {
	switch true {
	case update.Message != nil:
		return update.Message.Chat.ID
	case update.EditedMessage != nil:
		return update.EditedMessage.Chat.ID
	case update.ChannelPost != nil:
		return update.ChannelPost.Chat.ID
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost.Chat.ID
	}

	return -1
}

// extractMessageID get the Message.ID of an update (that is not inline-based)
func extractMessageID(update *echotron.Update) int {
	switch true {
	case update.Message != nil:
		return update.Message.ID
	case update.EditedMessage != nil:
		return update.EditedMessage.ID
	case update.ChannelPost != nil:
		return update.ChannelPost.ID
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost.ID
	case update.CallbackQuery != nil && update.CallbackQuery.Message != nil:
		return update.CallbackQuery.Message.ID
	}

	return -1
}

// extractMessageIDOpt generate and return the MessageIDOptions of a given update using the ID and SenderChat
func extractMessageIDOpt(update *echotron.Update) *echotron.MessageIDOptions {
	var (
		message *echotron.Message
		msgID   echotron.MessageIDOptions
		userID  int64
	)

	switch true {
	case update.Message != nil:
		message = update.Message
		userID = message.From.ID
	case update.EditedMessage != nil:
		message = update.EditedMessage
		userID = message.From.ID
	case update.ChannelPost != nil:
		message = update.ChannelPost
		userID = message.SenderChat.ID
	case update.EditedChannelPost != nil:
		message = update.EditedChannelPost
		userID = message.SenderChat.ID
	case update.InlineQuery != nil:
		msgID = echotron.NewInlineMessageID(update.InlineQuery.ID)
		return &msgID
	case update.ChosenInlineResult != nil:
		msgID = echotron.NewInlineMessageID(update.ChosenInlineResult.ResultID)
		return &msgID
	case update.CallbackQuery != nil:
		message = update.CallbackQuery.Message
		if message == nil {
			msgID = echotron.NewInlineMessageID(update.CallbackQuery.ID)
			return &msgID
		}
		userID = update.CallbackQuery.From.ID
	}

	if message == nil {
		return nil
	}
	msgID = echotron.NewMessageID(userID, message.ID)
	return &msgID
}

// extractCommand return the /command and the payload (other element separated by ' ' or '_' if /start)
func extractCommand(update *echotron.Update) (command string, payload []string) {
	var (
		text = extractText(update)
		ind  = strings.IndexRune(text, ' ')
	)

	if ind == -1 {
		return text, nil
	}

	command = text[:ind]
	payload = append(payload, text[ind+1:])
	if strings.ContainsRune(text[ind+1:], ' ') {
		payload = strings.Split(text[ind+1:], " ")
	} else if command == "/start" && strings.ContainsRune(text[ind+1:], '_') {
		payload = strings.Split(text[ind+1:], "_")
	}
	return
}

// extractName return the parsed FirstName of the user who sent the message
func extractName(update *echotron.Update) (FirstName string) {
	var user *echotron.User

	switch true {
	case update.Message != nil:
		user = update.Message.From
	case update.EditedMessage != nil:
		user = update.EditedMessage.From
	case update.ChannelPost != nil:
		user = update.ChannelPost.From
	case update.EditedChannelPost != nil:
		user = update.EditedChannelPost.From
	case update.InlineQuery != nil:
		user = update.InlineQuery.From
	case update.ChosenInlineResult != nil:
		user = update.ChosenInlineResult.From
	case update.CallbackQuery != nil:
		user = update.CallbackQuery.From
	}

	if user == nil {
		return "Unknown User"
	}
	FirstName = parseName(user.FirstName)
	if FirstName == "" {
		return "Unnamed User"
	}

	return
}

// parseName put the escaping sequence on name
func parseName(rawName string) string {
	var rx = regexp.MustCompile(`[\*\[\]\(\)\` + "`" + `~>#+\-=|{}.!]`)

	return rx.ReplaceAllString(rawName, "\\$0")
}

// Sender creates a CommandFunc that only returns the given message
func Sender(msg message.Any) CommandFunc {
	return func(bot *Bot, update *message.Update) message.Any {
		return msg
	}
}
