package robot

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// TOKEN is the Telegram API bot's token.
var TOKEN string

// LoadToken grabs the token from the command line arguments (os.Args) when next
// to executable name ex. `.\Parrbot.exe <TOKEN>`, or if token is inside a file
// and the name is followed by --readfrom and the file path, ex:
// .\DuelBot.exe --readfrom myfile.txt
func LoadToken() (err error) {
	var token string

	switch len(os.Args) {
	case 1:
		return errors.New("Missing TOKEN value")

	case 2:
		token = os.Args[1]

	case 3:
		if strings.ToUpper(os.Args[1]) != "--READFROM" {
			return errors.New("Invalid format")
		}
		content, err := os.ReadFile(os.Args[2])
		if err != nil {
			return err
		}
		token = strings.TrimSpace(string(content))

	default:
		return errors.New("Too many arguments")
	}

	err = validateToken(token)
	if err == nil {
		TOKEN = token
	}
	return
}

// validateToken tries to detect if a string is a vaild token
func validateToken(token string) error {
	var match, err = regexp.MatchString(`\d+:[\w\-]+`, token)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("Wrong format for TOKEN value")
	}

	return nil
}
