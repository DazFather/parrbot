package robot

import (
	"math"
	"regexp"

	"Parrbot/message"
)

type CommandFunc func(*Bot, *message.Update) message.Any

type Command struct {
	Name, Trigger string
	ReplyAt       message.UpdateType
	Scope         CommandFunc
}

var commands map[message.UpdateType]map[string]CommandFunc

func Divider(commandList []Command) (splitted map[message.UpdateType]map[string]CommandFunc) {
	splitted = make(map[message.UpdateType]map[string]CommandFunc, 0)

	for _, cmd := range commandList {
		for i := 0.0; i <= 6; i++ {
			t := message.UpdateType(math.Pow(2, i))
			if cmd.ReplyAt&t != 0 {
				if m := splitted[t]; m == nil {
					splitted[t] = make(map[string]CommandFunc, 0)
				}
				splitted[t][cmd.Trigger] = cmd.Scope
			}
		}
	}

	return
}

func Select(update *message.Update) CommandFunc {
	var (
		trigger string
		filter  message.UpdateType
	)

	switch true {
	case update.Message != nil:
		if update.Message.Text != "" {
			trigger = extractTrigger(update.Message.Text)
		} else {
			trigger = extractTrigger(update.Message.Caption)
		}
		filter = message.MESSAGE
	case update.EditedMessage != nil:
		trigger = extractTrigger(update.EditedMessage.Text)
		filter = message.EDITED_MESSAGE
	case update.ChannelPost != nil:
		trigger = extractTrigger(update.ChannelPost.Text)
		filter = message.CHANNEL_POST
	case update.EditedChannelPost != nil:
		trigger = extractTrigger(update.EditedChannelPost.Text)
		filter = message.EDITED_CHANNEL_POST
	case update.CallbackQuery != nil:
		trigger = extractTrigger(update.CallbackQuery.Data)
		filter = message.CALLBACK_QUERY
	}

	return commands[filter][trigger]
}

func extractTrigger(caption string) string {
	var rgxp = regexp.MustCompile(`^/\w+`)
	return rgxp.FindString(caption)
}

func LoadCommands(commandList []Command) {
	commands = make(map[message.UpdateType]map[string]CommandFunc, 0)
	commands = Divider(commandList)
}
