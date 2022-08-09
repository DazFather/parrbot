package robot

import (
	"log"
	"regexp"

	"github.com/DazFather/parrbot/message"

	"github.com/NicoNex/echotron/v3"
)

// Command is a bot's command declaration that compose the command list
type Command struct {
	Description string             // A description of the command that will be displayed on the "/" menu if the ReplyAt includes MESSAGE
	Trigger     string             // Needs to start with the '/' character. Is the string that if contained at the start of the update would run the Scope
	ReplyAt     message.UpdateType // Tells witch UpdateType(s) the bot will reply at, sum them to put more
	CallFunc    CommandFunc        // The actual function that the bot will run
}

// CommandFunc is a custom type that rapresent a command that the bot should be able to run
type CommandFunc func(*Bot, *message.Update) message.Any

// commands is where all the command will be stored
var commands map[message.UpdateType]map[string]CommandFunc

// Divider divide the command list and cast it in a form that is more efficenct
func Divider(commandList []Command) (splitted map[message.UpdateType]map[string]CommandFunc) {
	splitted = make(map[message.UpdateType]map[string]CommandFunc, 0)

	var cmdMenu []echotron.BotCommand

	for _, cmd := range commandList {

		if cmd.ReplyAt&message.MESSAGE != 0 && cmd.Trigger != "" && cmd.Description != "" {
			cmdMenu = append(cmdMenu, echotron.BotCommand{
				Command:     cmd.Trigger,
				Description: cmd.Description,
			})
		}

		for i := 0; i <= 9; i++ {
			t := message.UpdateType(1 << i)
			if cmd.ReplyAt&t != 0 {
				if m := splitted[t]; m == nil {
					splitted[t] = make(map[string]CommandFunc, 0)
				}
				splitted[t][cmd.Trigger] = cmd.CallFunc
			}
		}
	}

	if len(cmdMenu) == 0 {
		return
	}

	res, err := message.GetAPI().SetMyCommands(nil, cmdMenu...)
	if err != nil {
		log.Fatal("SetMyCommands error: ", err)
	}
	if res.Result != true || res.Ok != true {
		log.Fatal("SetMyCommands wrong response: ", res)
	}

	return
}

// Select take an update and verify it's type and then trigger in order to return the appropriate function (or nil)
func Select(update *message.Update) CommandFunc {
	var (
		trigger string
		filter  message.UpdateType
	)

	switch true {
	case update.Message != nil:
		trigger = extractTrigger(update.Message.Text)
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
	case update.InlineQuery != nil:
		filter = message.INLINE_QUERY
	case update.CallbackQuery != nil:
		trigger = extractTrigger(update.CallbackQuery.Data)
		filter = message.CALLBACK_QUERY
	case update.MyChatMember != nil:
		filter = message.MY_CHAT_MEMBER
	case update.ChatMember != nil:
		filter = message.CHAT_MEMBER
	case update.ChatJoinRequest != nil:
		filter = message.CHAT_JOIN_REQUEST
	}

	return commands[filter][trigger]
}

// extractTrigger is in charge of extracting the trigger from a caption originated by the update
func extractTrigger(caption string) string {
	var rgxp = regexp.MustCompile(`^/\w+`)
	return rgxp.FindString(caption)
}

// LoadCommands saves the given commandList in a form that is more efficenct
func LoadCommands(commandList []Command) {
	commands = make(map[message.UpdateType]map[string]CommandFunc, 0)
	commands = Divider(commandList)
}
