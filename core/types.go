package core

type GameContext struct {
	CurrentPlayer Player
	CurrentPhase  string
	RoundCount    int
	TurnOrder     []Player
}

type PhaseHandler struct {
	CurrentPhase string
}

type Player struct {
	ID    string
	Name  string
	Score int
}

type Rule struct {
	SignalType string
	Condition  func(GameState, GameContext) bool
	Action     func(GameState, GameContext) (Action, error)
}
