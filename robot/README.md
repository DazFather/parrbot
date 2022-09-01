## robot package

This package contains core functionality for building your bot.

### Create a bot
To create and start a bot the only function needed is `Start` that you can find
on this package.
This function will load the configuration, retrieve the token and save given commands
and create the menu on the chat.
> For this reason if you need to change your configuration do it before you call this function

As previously mentioned this function will also allow to set your commands. There are some important
notions to keep in mind when you create a command:
If present the _Trigger_ MUST start with a "_/_". When empty or not given the command will reply at every updates of all the types included in the _ReplyAt_ field.
The field _Description_, if present, will generate an actual description of the command in the menu usable inside on the chat but only if _ReplyAt_ includes also `message.MESSAGE`.

### API Token
The Telegram Bot API TOKEN is normally given in input as a program argument of your application like this:  $`<EXECUTABLE> <TOKEN>`

If you don't want to insert your API TOKEN in the command line all the time you can save it on a file
(the file must contains only the bot token) and you can use the _readfrom_ command followed by the path of your file, like this: $`<EXECUTABLE> --readfrom <PATH>`
This command is case insensitive and can work with none, one or two "_-_".

There is also the possibility to inserting the TOKEN directly in your application using the configuration.
To do so just use the variable `Config` already present in this package and use the `SetAPIToken` method.
All configurations will be loaded on the Start command so edit the configurations (including the mentioned method) before.

---

> _Part of the [Parr(B)ot](https://github.com/DazFather/parrbot) framework._
