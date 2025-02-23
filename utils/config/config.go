package config

// GameConfig defines a game's configuration
type GameConfig struct {
	Players       int
	WinConditions []string
}

// Loader loads configurations from files
type Loader interface {
	Load(path string) (*GameConfig, error)
}
