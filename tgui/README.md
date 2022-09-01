## tgui package

"_tgui_" stands for _TeleGram User Interface_. This package contains functionality and utilities for build or use user interfaces like menu or keyboards utilities.
It is not a core package, so is NOT necessary for the correct execution of a bot.

### Main features
These are the main features of this package:

- **ShowMessage** is a function that will display a message to the user by modifying the previous one when called via callback or sending a new one otherwhise.

- **Menu** is an interface with currently two implementation:
    > **PagedMenu** that allows to create an inline menu that allows user to navigate between previous and next page
    > **InlineMenu** that allows to create more complex and nested inline menus

all the pages of the menu are functions that allows to show contents in a dynamic way

- **Shorter type alias** like EditOptions _(echotron.MessageTextOptions)_, InlineButton _(echotron.InlineKeyboardButton)_ or KeyButton _(echotron.KeyboardButton)_

- **Utilities** for building and rearranging keyboards, managing options to edit messages or pages or create parrbot commands



---

> _Part of the [Parr(B)ot](https://github.com/DazFather/parrbot) framework._
