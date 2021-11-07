package robot

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
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

/* GrabToken grab the token from using the command line arguments (os.Args)
 * put the token next to the executable file name on the console or
 * use put readfrom followed by the file in witch there is the token as example:
 * .\DuelBot.exe --readfrom mytoken.txt
 */
func GrabToken() (token string, err error) {
	switch len(os.Args) {
	case 1:
		return "", errors.New("Missing TOKEN value")

	case 2:
		token = os.Args[1]

	case 3:
		if strings.ToUpper(os.Args[1]) != "--READFROM" {
			return "", errors.New("Invalid format")
		}
		if content, err := os.ReadFile(os.Args[2]); err != nil {
			return "", err
		} else {
			token = strings.TrimSpace(string(content))
		}

	default:
		return "", errors.New("Too many arguments")
	}

	err = validateToken(token)
	return
}

// validateToken tries to detect if a string is a vaild token
func validateToken(token string) error {
	match, err := regexp.MatchString(`\d+:[\w\-]+`, token)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("Wrong format for TOKEN value")
	}

	return nil
}
