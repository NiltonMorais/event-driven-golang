package handler

import (
	"context"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type Handler interface {
	Execute(ctx context.Context, e event.DomainEvent) error
}
