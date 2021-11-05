package robot

import (
	"log"
)

// TOKEN is the Telegram API bot's token.
var TOKEN string

func LoadToken() {
	if rawToken, err := GrabToken(); err != nil {
		log.Fatal(err)
	} else {
		TOKEN = rawToken
	}
}
