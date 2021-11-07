package robot

import (
	"log"
)

// TOKEN is the Telegram API bot's token.
var TOKEN string

// LoadToken get the Telegram API bot's token and put it into the global variable "TOKEN"
func LoadToken() {
	if rawToken, err := GrabToken(); err != nil {
		log.Fatal(err)
	} else {
		TOKEN = rawToken
	}
}
