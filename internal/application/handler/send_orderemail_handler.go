package handler

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type SendOrderEmailHandler struct {
}

func NewSendOrderEmailHandler() *SendOrderEmailHandler {
	return &SendOrderEmailHandler{}
}

func (h *SendOrderEmailHandler) Execute(ctx context.Context, e event.DomainEvent) {
	payload := event.GetPayload[event.OrderCreatedEvent](e)
	fmt.Println("--- SendOrderEmailHandler ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.Order.GetTotalPrice())
}
