# Tabulludo

Tabulludo is a modular boardgame framework written in Go. It provides a set of standardized interfaces and components for managing game state, handling actions and events, enforcing game rules, and managing game lifecycle processes. It is intended to be used as a base for developing boardgames rather than a standalone application. It is currently under development and not all features are available.

## Features

- **Game State Management:**  
  Define and manage game state using the `GameState` interface. Components can be dynamically added and retrieved through a standardized component system.

- **Action System:**  
  Implement game logic through actions that are validated and executed against the game state. Actions can generate events to notify external systems of state changes.

- **Rule Engine:**  
  Process game signals through a modular rule engine. Rules combine signal handling with conditional logic to trigger appropriate game actions based on the current state and context.

- **Turn Management:**  
  Utilize the `TurnManager` interface to manage game turns and phases, supporting dynamic turn ordering and phase transitions.

- **Lifecycle Control:**  
  Manage game sessions using lifecycle methods (start, pause, stop, save, load) provided by the `GameLifecycle` interface.

- **External Communication:**  
  Use the `EventBus` system to notify external systems (UI, network, etc.) about game state changes. Events provide a clean separation between core game logic and external representations.

- **Basic Configuration System:**  
  Define game parameters through a configuration system supporting player counts and win conditions, with extensibility for additional settings.

## Architecture Overview

The project is organized into several key packages:

- **Core (`core/`):**  
  Contains the fundamental interfaces and types for game components.
  - **interfaces.go:** Defines core interfaces including `GameState`, `Action`, `Effect`, `GameSignal`, and `RuleEngine`.
  - **lifecycle.go:** Declares the `GameLifecycle` interface for game session control.
  - **container.go:** Provides the `GameContainer` for organizing game state, context, and dependencies.
  - **types.go:** Defines common types like `GameContext`, `Player`, `PhaseHandler`, and `Rule`.
  - **error.go:** Defines standard error types for common game operation failures.

- **Events (`utils/events/`):**  
  Implements an event bus system for communicating game changes to external systems like UI and network components.

- **Configuration (`utils/config/`):**  
  Defines structures and interfaces for loading game configurations.

- **State Management (`utils/state/`):**  
  Offers functionality for diffing game states and generating patches to track changes or support undo/redo mechanisms.

## Getting Started

Since Tabulludo is designed to be used as a dependency, you can easily add it to your project using Go modules. Simply run:

```bash
go get github.com/psychometry/tabulludo
```

*🚧 Demo implementations are currently in development and will be added to the `examples/` directory once the core framework stabilizes.*

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
