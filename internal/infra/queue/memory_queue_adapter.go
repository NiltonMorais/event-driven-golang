package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type MemoryQueueAdapter struct {
	listeners map[string][]Listener
}

func NewMemoryQueueAdapter() *MemoryQueueAdapter {
	return &MemoryQueueAdapter{
		listeners: make(map[string][]Listener),
	}
}

func (eb *MemoryQueueAdapter) ListenerRegister(eventType reflect.Type, handler func(w http.ResponseWriter, r *http.Request)) {
	eb.listeners[eventType.Name()] = append(eb.listeners[eventType.Name()], Listener{eventType, handler})
}

func (eb *MemoryQueueAdapter) Publish(ctx context.Context, eventPayload interface{}) error {
	eventType := reflect.TypeOf(eventPayload)
	payloadJson, _ := json.Marshal(eventPayload)

	log.Printf("--- Publish %s ---", eventType)

	for _, listener := range eb.listeners[eventType.Name()] {
		w := NewQueueResponseWriter()
		body := bytes.NewBuffer(payloadJson)
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, eventType.Name(), body)
		if err != nil {
			return err
		}

		listener.callback(w, r)
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

func (eb *MemoryQueueAdapter) StartConsuming(ctx context.Context, queueName string) error {
	log.Printf("--- MemoryQueueAdapter StartConsuming queue %s ---", queueName)
	return nil
}
