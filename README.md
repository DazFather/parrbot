# Parr(B)ot framework

[![Go Report Card](https://goreportcard.com/badge/github.com/DazFather/Parrbot)](https://goreportcard.com/report/github.com/DazFather/Parrbot)
[![Telegram](https://img.shields.io/badge/Parr(B)ot%20News-blue?logo=telegram&style=flat)](https://t.me/+3_LBajtkqUgzOTFk) 

A just born Telegram bot framework in [Go](https://golang.org) built on top of the [echotron library](https://github.com/NicoNex/echotron).
> You can call it Parrot, Parr-Bot, Parrot Bot, is up to you


## Philosophy and Design

**Framework** not library: Parr(B)ot main focus is to help you onbetter organize your code, in a way that is highly scalable and reusable.
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

## Usage

 _Step 1._ - Clone this repo (and be sure to have Go installed)

 _Step 2._ - Create your own awesome bot (see an example on main.go)

 _Step 3._ - Make it fly!
