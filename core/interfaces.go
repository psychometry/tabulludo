package core

//State and Context govern the general flow of the game.

type GameState interface {
	// Component management
	GetComponent(id string) (interface{}, error)        // Get a single component by ID
	AddComponent(component interface{}) (string, error) // Add a new component, returns ID
	RemoveComponent(id string) error                    // Remove a component by ID

	// Collection management
	CreateCollection(name string, componentType string) error             // Create a new empty collection
	AddToCollection(collectionName string, componentID string) error      // Add component to collection
	RemoveFromCollection(collectionName string, componentID string) error // Remove component from collection
	GetCollection(collectionName string) ([]interface{}, error)           // Get all components in a collection

	// State tracking
	GetPlayerState(playerID string) (map[string]interface{}, error)        // Get player-specific state
	UpdatePlayerState(playerID string, state map[string]interface{}) error // Update player state
}

//Actions are the individual moves that can be made, direct
type Action interface {
	Validate(state GameState) error
	Execute(state GameState) ([]Event, error)
}

//Effect are Changes made to the context
type Effect interface {
	Apply(context GameContext) error
}

// Signal are the events that trigger the rules engine
type GameSignal interface {
	Type() string
	Data() interface{}
}

//Rules validate and execute actions, and check win conditions
type RuleEngine interface {
	AddRule(rule Rule)
	ProcessSignal(signal GameSignal, state GameState, context GameContext, depth int) error
}

type TurnManager interface {
	NextPlayer() Player          // Returns the next player to act
	CurrentPhase() string        // Identifies the current phase (e.g., "main", "combat")
	AdvancePhase() error         // Progress to next phase
	TurnOrder() []Player         // Full turn sequence
	SetTurnOrder([]Player) error // Dynamic turn order adjustment
	RoundCount() int             // Track game progress
}

// ComponentFactory is responsible for creating and managing game components
type ComponentFactory interface {
	// CreateComponent creates a new component of the specified type
	CreateComponent(componentType string) (interface{}, error)

	// RegisterComponent registers a component type with its constructor
	RegisterComponent(componentType string, constructor func() interface{}) error
}

// Event represents a game event that has occurred, meant for communication to external systems
type Event interface {
	// Type returns the event type identifier
	Type() string
	// Data returns the event payload
	Data() interface{}
}

type EventBus interface {
	Subscribe(eventType string, handler EventHandler)
	Publish(event Event)
}

// EventHandler is a function type that handles game events
type EventHandler func(Event)
