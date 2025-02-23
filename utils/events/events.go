package events

import (
	"Tabulludo/core"
	"sync"
)

// EventBus implements core.EventBus
type EventBus struct {
	listeners map[string][]core.EventHandler
	mu        sync.RWMutex
}

// NewEventBus creates a new event bus instance
func NewEventBus() core.EventBus {
	return &EventBus{
		listeners: make(map[string][]core.EventHandler),
	}
}

// Subscribe registers an event handler for a specific event type
func (eb *EventBus) Subscribe(eventType string, handler core.EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners[eventType] = append(eb.listeners[eventType], handler)
}

// Publish broadcasts an event to all registered handlers of that event type
func (eb *EventBus) Publish(event core.Event) {
	eb.mu.RLock()
	handlers, exists := eb.listeners[event.Type()]
	eb.mu.RUnlock()

	if !exists {
		return
	}

	for _, handler := range handlers {
		go handler(event)
	}
}
