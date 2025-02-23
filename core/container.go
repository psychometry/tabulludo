package core

type GameContainer struct {
	state        GameState
	context      GameContext
	dependencies GameDependencies
}

type GameDependencies struct {
	StateFactory     func() GameState
	ContextFactory   func() GameContext
	RuleEngine       RuleEngine
	ComponentFactory ComponentFactory
	EventBus         EventBus
}

func NewGameContainer(deps GameDependencies) *GameContainer {
	return &GameContainer{
		state:        deps.StateFactory(),
		context:      deps.ContextFactory(),
		dependencies: deps,
	}
}
