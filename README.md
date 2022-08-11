# Parr(B)ot framework

[![Go Report Card](https://goreportcard.com/badge/github.com/DazFather/Parrbot)](https://goreportcard.com/report/github.com/DazFather/Parrbot)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazFather/parrbot.svg)](https://pkg.go.dev/github.com/DazFather/parrbot)
[![Telegram](https://img.shields.io/badge/Parr(B)ot%20News-blue?logo=telegram&style=flat)](https://t.me/+3_LBajtkqUgzOTFk)

A just born Telegram bot framework in [Go](https://golang.org) built on top of the [echotron library](https://github.com/NicoNex/echotron).
> You can call it Parrot, Parr-Bot, Parrot Bot, is up to you


## Philosophy and Design

**Framework** not library: Parr(B)ot main focus is to help you to better organize your code in a way that is highly scalable and reusable.
Libraries are cool but if not used carefully tends to transform your code into an extremely difficult to debug spaghetti-code monster.
> _Framework is like the nest of your project_ 🦜

**Modular**: Modules helps developer to access, reuse and expand their codes and also give a sense of order to your project that will help in the long run. The base structure of Parr(b)ot is design to be easy to maintain in the future
> _Put together some code and a pair of wings to build your own robo-parrot_ 🦜

**Typed**: We love types! They allows you to write solid codes and if the language is (like Go) compiled you will not need to do any runtime checking of the entry values.
That said sometimes having access to many types and constructor can be overwhelming this is the reason why the interface can be very handy to manage different data-structure that have similar behavior.
You can find an example of what we mean on the "incoming.go" file (inside the "message" module).
> _We think that biodiversity is great_ 🦜

**Customizable**: The parr(b)ot modules are not made to be untouched but the exact opposite. User is more then welcome to change the framework behavior in order for example to extend a module with more useful functions or types, to change the information carried by the Bot or to initialize the bot in a different way. Is up to you.
The framework will try to guide you using comments and naming to go at the level of deepness that you desire
> _Explore the jungle, dangerous but fun_ 🦜

### Main features

_Okay, cool but what it does actually mean? How does all of this translate in something useful for me? Why should I use Parr(B)ot?_
It's time to talk about the main features of this framework:
 - **Linear design**: each command in Parr(B)ot can return a message that will be automatically sent. Althow you can still return nil and send a message meanwhile, this type of approach wants to invite user to have a function that returns a message when needed, in this way if something goes wrong (message is wrong or not appear) is easier to identify and locate the bug and makes your code more organized
 - **Edit-on-sync**: when sending a request to Telegram to edit an message, if everything will go right the same edits will be applied on the message variable that you used to send the request too, to reduce the instantiation of new variables and help you to keep tracks of the changes without loosing anything on the way
 - **Ready to go**: Parr(B)ot helps user to quickly creates Telegram idiomatic commands, by encouraging usage of triggers that start with a '/' and filling the command description right away, on top of that the framework will also care about managing the Telegram Bot API TOKEN, without any need to hardcode it and difficulties when making the code public.
 - **Here to help**: lots of utilities to help you achive what you want faster, without warring about common stuffs like creating a menu system, run your /commands when they are supposed to, managing some buttons. Parr(B)ot already did all these things for you

_...and more are coming_

## Documentation

[Here](https://pkg.go.dev/github.com/DazFather/parrbot) there is the official documentation of parrbot. As you will see is divided in 3 main packages / directories:
 - message - (Core) manage incoming / outgoing message-related stuffs
 - robot - (Core) manage bots sessions and commands
 - tgui - toolkit for user interfaces like menus or keyboards utilities

Parr(B)ot makes massively use of the [echotron library](https://pkg.go.dev/github.com/NicoNex/echotron/v3), it might be useful also it's doc. Keep in mind that the echotron library is almost 1:1 with the Telegram's Bot API, if you are unsure about the meaning of certain fields you can always have a look to the [Telegram's doc](https://core.telegram.org/bots/api).

## Usage

 - **Step 1. Import or clone Parr(B)ot** - There are two main ways to use Parr(B)ot: cloning the repo or via _import_ of the needed package simply like this: `import "github.com/DazFather/parrbot/<package>"`. Notice that "`message`" and "`robot`" packages contains core functionality.
> Make sure you have Go installed. You can check the required version on the [go.mod](./go.mod) file (keep [this](https://go.dev/doc/go1compat) in mind). The Go's official [guide for managing dependency](https://go.dev/doc/modules/managing-dependencies) might be useful too

 - **Step 2. Create your own awesome bot** - Use the function robot.Start and fill it with a list of robot.Command that the bot will execute when triggered
> Check out [main.go](./main.go) for a usage example

 - **Step 3. Make it fly!** - After building your project, run the bot and use as argument the API TOKEN, or save it on a _".txt"_ file and use `--readfrom` followed by the file path.
> You can get the Telegram's Bot API TOKEN by creating a bot using [@BotFather](https://t.me/BotFather)
