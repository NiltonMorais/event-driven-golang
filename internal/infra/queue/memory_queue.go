package queue

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type MemoryQueue struct {
	listeners map[string][]Listener
}

type Listener struct {
	eventType reflect.Type
	callback  func(ctx context.Context, e event.DomainEvent)
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		listeners: make(map[string][]Listener),
	}
}

func (eb *MemoryQueue) Register(eventType reflect.Type, callback func(ctx context.Context, e event.DomainEvent)) {
	eb.listeners[eventType.Name()] = append(eb.listeners[eventType.Name()], Listener{eventType, callback})
}

func (eb *MemoryQueue) Publish(ctx context.Context, eventPayload interface{}) {
	eventType := reflect.TypeOf(eventPayload)
	domainEvent := event.DomainEvent{
		Type:    eventType,
		Date:    time.Now(),
		Payload: eventPayload,
	}

	log.Printf("--- Publish %s ---", eventType)
	// log.Printf("Data: %s", eventPayload)

	for _, listener := range eb.listeners[eventType.Name()] {
		listener.callback(ctx, domainEvent)
	}
}
