package message

import (
	"github.com/NicoNex/echotron/v3"
)

type Update echotron.Update

type UpdateType uint8

const (
	MESSAGE              UpdateType = 1  // 0000001
	EDITED_MESSAGE                  = 2  // 0000010
	CHANNEL_POST                    = 4  // 0000100
	EDITED_CHANNEL_POST             = 8  // 0001000
	INLINE_QUERY                    = 16 // 0010000
	CHOSEN_INLINE_RESULT            = 32 // 0100000
	CALLBACK_QUERY                  = 64 // 1000000
)

const ANY UpdateType = 127 // 1111111
