package queue

import (
	"context"
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type Listener struct {
	eventType reflect.Type
	callback  func(ctx context.Context, e event.DomainEvent) error
}
