package queue

import (
	"context"
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type Queue interface {
	Register(eventType reflect.Type, callback func(ctx context.Context, e event.DomainEvent))
	Publish(ctx context.Context, eventPayload interface{})
}
