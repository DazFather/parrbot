package robot

import (
	"time"

	"github.com/DazFather/parrbot/message"
)

// Config contains all the default Parrbot configurations. Edit them before robot.Start
var Config = ParrbotConfig{
	DeleteSessionTimer: time.Hour * 2,
	// by default token will be loaded using os.Args
}

// ParrbotConfig defines all the possible configurations of your parr-bot
type ParrbotConfig struct {
	DeleteSessionTimer time.Duration // time after witch the bot session will self distruct by the dispatcher
	token              string        // Telegram API bot's token.
}

// KeepActiveSessions sets the DeleteSessionTimer = 0 causing all session to stay active
func (c *ParrbotConfig) KeepActiveSessions() {
	c.DeleteSessionTimer = 0
}

// SetAPIToken sets the Telegram Bot API token to the given one if valid
func (c *ParrbotConfig) SetAPIToken(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}
	c.token = token
	message.LoadAPI(token)
	return nil
}

// loadDefaultToken set the token to by checking on os.Args and validate result
func (c *ParrbotConfig) loadDefaultToken() error {
	var token, err = retrieveToken()
	if err == nil {
		err = c.SetAPIToken(token)
	}
	return err
}

// init is used to initialize a ParrbotConfig to
func (c *ParrbotConfig) init() error {
	// if token is un-initilized load default
	if c.token == "" {
		return c.loadDefaultToken()
	}

	return nil
}
