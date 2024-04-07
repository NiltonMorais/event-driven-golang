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

func (h *SendOrderEmailHandler) Execute(ctx context.Context, e event.DomainEvent) error {
	payload, err := event.GetPayload[event.OrderCreatedEvent](e)
	if err != nil {
		return err
	}
	fmt.Println("--- SendOrderEmailHandler ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
