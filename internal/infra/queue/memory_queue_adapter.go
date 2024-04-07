package queue

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type MemoryQueueAdapter struct {
	listeners map[string][]Listener
}

func NewMemoryQueueAdapter() *MemoryQueueAdapter {
	return &MemoryQueueAdapter{
		listeners: make(map[string][]Listener),
	}
}

func (eb *MemoryQueueAdapter) ListenerRegister(eventType reflect.Type, callback func(ctx context.Context, e event.DomainEvent) error) {
	eb.listeners[eventType.Name()] = append(eb.listeners[eventType.Name()], Listener{eventType, callback})
}

func (eb *MemoryQueueAdapter) Publish(ctx context.Context, eventPayload interface{}) error {
	eventType := reflect.TypeOf(eventPayload)
	payloadJson, _ := json.Marshal(eventPayload)
	domainEvent := event.DomainEvent{
		Type:    eventType,
		Date:    time.Now(),
		Payload: payloadJson,
	}

	log.Printf("--- Publish %s ---", eventType)
	// log.Printf("Data: %s", eventPayload)

	for _, listener := range eb.listeners[eventType.Name()] {
		err := listener.callback(ctx, domainEvent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (eb *MemoryQueueAdapter) Connect(ctx context.Context) error {
	log.Println("--- MemoryQueueAdapter connected ---")
	return nil
}

func (eb *MemoryQueueAdapter) Disconnect(ctx context.Context) error {
	log.Println("--- MemoryQueueAdapter disconnected ---")
	return nil
}

func (eb *MemoryQueueAdapter) StartConsuming(ctx context.Context) error {
	log.Println("--- MemoryQueueAdapter StartConsuming ---")
	return nil
}
