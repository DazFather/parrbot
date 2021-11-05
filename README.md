# Parr(B)ot framework

A just born Telegram bot framework in [Go](https://golang.org) based on top of the [echotron library](https://github.com/NicoNex/echotron).
> You can call it Parrot, Parr-Bot, Parrot Bot, is up to you


## Philosophy and Design

**Framework vs Library**: Let's start to define what Parr(B)ot is not: A library.
A library is a collection of function that you can use _as you want*_ in your code, without a specific order (es. [Pandas](https://pandas.pydata.org/), [Bootstrap](https://getbootstrap.com/)).
A framework on the other hand gives you order for both files and code structure (es. [Django](https://www.djangoproject.com/), [Angular](https://angular.io/)) and tends to have more layer of abstraction.
> _Framework is like the nest of your project_ 🦜

**Modular**: Modules helps developer to access, reuse and expand their codes and also give a sense of order to your project that will help in the long run. The base structure of Parr(b)ot is design to be easy to maintain in the future
> _Makes together some part and build your own robo-parrot_ 🦜

**Typed**: We love types! They allows you to write solid codes and if the language is (like Go) compiled you will not need to do any runtime checking of the entry values.
That said sometimes having access to many types and constructor can be overwhelming this is the reason why the interface can be very handy to manage different data-structure that have similar behavior.
You can find an example of what we mean on the "incoming.go" file (inside the "message" module).
> _We think that biodiversity is great_ 🦜

**Customizable**: The parr(b)ot modules are not made to be untouched but the exact opposite. User is more then welcome to change the framework behavior in order for example to extend a module with more useful functions or types, to change the information carried by the Bot or to initialize the bot in a different way. Is up to you.
The framework will try to guide you using comments and naming to go at the level of deepness that you desire
> _Explore the jungle, is fun_ 🦜

## Usage

 _Step 1._ - Clone this repo (and be sure to have Go installed)

 _Step 2._ - Create your own awesome bot (see an example on main.go)
