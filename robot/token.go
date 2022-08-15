package robot

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// retrieveToken grabs the token from the command line arguments (os.Args) when
// next to executable name ex. `.\Parrbot.exe <TOKEN>`, or if token is inside a
// file and the name is followed by --readfrom and the file path, ex:
// .\DuelBot.exe --readfrom myfile.txt
func retrieveToken() (token string, err error) {
	switch len(os.Args) {
	case 1:
		err = errors.New("Missing TOKEN value")

	case 2:
		token = os.Args[1]

	case 3:
		if !regexp.MustCompile(`(?i)-{0,2}readfrom`).MatchString(os.Args[1]) {
			return "", errors.New("Invalid format")
		}

		content, err := os.ReadFile(os.Args[2])
		if err == nil {
			token = strings.TrimSpace(string(content))
		}

	default:
		err = errors.New("Too many arguments")
	}

	return
}

// validateToken tries to detect if a string is a vaild token
func validateToken(token string) error {
	if !regexp.MustCompile(`\d+:[\w\-]+`).MatchString(token) {
		return errors.New("Wrong format for TOKEN value")
	}
	return nil
}
