package core

import "errors"

var (
	ErrInvalidAction = errors.New("invalid action")
	ErrGameOver      = errors.New("game is over")
	ErrStateCorrupt  = errors.New("game state is corrupt")
)
