package state

import "github.com/psychometry/tabulludo/core"

// Differ calculates differences between game states
type Differ interface {
	Diff(a, b core.GameState) []Patch
}

// Patch represents a state change
type Patch struct {
	Path  string
	Value interface{}
}
