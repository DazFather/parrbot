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

// divide the command list and cast it in a form that is more efficenct
func divide(commandList []Command) (splitted map[message.UpdateType]map[string]CommandFunc) {
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

	res, err := message.API().SetMyCommands(nil, cmdMenu...)
	if err != nil {
		log.Fatal("SetMyCommands error: ", err)
	}
	if res.Result != true || res.Ok != true {
		log.Fatal("SetMyCommands wrong response: ", res)
	}

	return
}

// Select take an update and verify it's type and then trigger in order to return
// the appropriate function (or nil). If robot.Start is used (as racommanded),
// probably, there is no need to use this function
func Select(update *message.Update) CommandFunc {
	var (
		rgx     = regexp.MustCompile(`^/\w+`)
		trigger string
		filter  message.UpdateType
	)

	switch true {
	case update.Message != nil:
		trigger = rgx.FindString(update.Message.Text)
		filter = message.MESSAGE
	case update.EditedMessage != nil:
		trigger = rgx.FindString(update.EditedMessage.Text)
		filter = message.EDITED_MESSAGE
	case update.ChannelPost != nil:
		trigger = rgx.FindString(update.ChannelPost.Text)
		filter = message.CHANNEL_POST
	case update.EditedChannelPost != nil:
		trigger = rgx.FindString(update.EditedChannelPost.Text)
		filter = message.EDITED_CHANNEL_POST
	case update.InlineQuery != nil:
		filter = message.INLINE_QUERY
	case update.ChosenInlineResult != nil:
		filter = message.CHOSEN_INLINE_RESULT
	case update.CallbackQuery != nil:
		trigger = rgx.FindString(update.CallbackQuery.Data)
		filter = message.CALLBACK_QUERY
	case update.ShippingQuery != nil:
		filter = message.SHIPPING_QUERY
	case update.PreCheckoutQuery != nil:
		filter = message.PRE_CHECKOUT_QUERY
	case update.MyChatMember != nil:
		filter = message.MY_CHAT_MEMBER
	case update.ChatMember != nil:
		filter = message.CHAT_MEMBER
	case update.ChatJoinRequest != nil:
		filter = message.CHAT_JOIN_REQUEST
	}

	return commands[filter][trigger]
}

// LoadCommands saves the given commandList in a form that is more efficenct for
// the bot to retrive. Use this function one time only, is necessary for Select
// to work. If robot.Start is used (as racommanded), probably, there is no need
// to use this function
func LoadCommands(commandList []Command) {
	commands = make(map[message.UpdateType]map[string]CommandFunc, 0)
	commands = divide(commandList)
}
