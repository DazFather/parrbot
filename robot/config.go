package robot

import "time"

// Config contains all the default Parrbot configurations. Edit them before robot.Start
var Config = ParrbotConfig {
    DeleteSessionTimer: time.Hour * 2,
}

// ParrbotConfig defines all the possible configurations of your parr-bot
type ParrbotConfig struct {
    DeleteSessionTimer time.Duration // time after witch the bot session will self distruct by the dispatcher
}

// KeepActiveSessions sets the DeleteSessionTimer = 0 causing all session to stay active
func (c *ParrbotConfig) KeepActiveSessions() {
    c.DeleteSessionTimer = 0
}
