package queue

import (
	"context"
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type Queue interface {
	ListenerRegister(eventType reflect.Type, callback func(ctx context.Context, e event.DomainEvent) error)
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Publish(ctx context.Context, eventPayload interface{}) error
	StartConsuming(ctx context.Context, queueName string) error
}
