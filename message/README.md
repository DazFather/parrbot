## message package

This package contains core functionality for interacting with Telegram, managing updates and messages.

### Handling updates
An `Update` is represent any input received. Normally only one of the field (not including the ID) will be populated with the infos received from Telegram.
An `UpdateType` indicates what field is being populated ans is used on the _robot_ package to specify to what update a command needs to reply on, in this context you can also use the sum operator '+' or the bitwise or operator '|' to include multiple types

### Managing message
**Incoming** and **sent** messages are represented as `UpdateMessage` witch are very similar to the Telegram representation but present some wrappers like _Forward_, _Media_ and _SystemNotification_ that holds a specific group of informations.
Due to Telegram's limitation is only possible to edit a message sent by the bot itself (except for messages published in channels), but is still possible to delete
any type of messages (sent in the last 24h). When editing a message, the edit will be synchronized, this means that if returned error is nil, the variable that is being used to edit a message will change accordingly.
There is also the possibility to create a **reference** of the sent message thanks to the `NewReference` function. This struct is way lighter and is capable of both editing and deleting the associated message without the sync capability.

**Outgoing** messages are represented by different structs depending by the type of message that we are sending (_Text_ for plain text messages, _Photo_ for messages that contain picture, _Sticker_ ...). Depending on the type of message you can use various methods like ClipInlineKeyboard to set some specific options. All messages implements the `Any` interface thanks to the `Send` method.

### Echotron interoperability
Parr(B)ot and in particular this package makes an extensive use of the Echotron library. This means that sometimes user will need to deal with some echotron's data structure.

Normally the main logic of your bot when using Parr(B)ot is delegated to the package robot, but this package makes available functions that allows developer to use it without the need of other parrbot's packages but only using echotron (even though is still not recommended):
- `CastUpdate` to allow conversion from an echotron.Update into an Update as is re-defined in this library
- `LoadAPI` and to save the API TOKEN externally, this will make methods contained in this package like Send or the edits works.
- `API` to retrieve the echotron.API it when needed.

---

> _Part of the [Parr(B)ot](https://github.com/DazFather/parrbot) framework._
