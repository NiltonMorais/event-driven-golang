package handler

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type ProcessOrderPaymentHandler struct {
}

func NewProcessOrderPaymentHandler() *ProcessOrderPaymentHandler {
	return &ProcessOrderPaymentHandler{}
}

func (h *ProcessOrderPaymentHandler) Execute(ctx context.Context, e event.DomainEvent) error {
	payload, err := event.GetPayload[event.OrderCreatedEvent](e)
	if err != nil {
		return err
	}
	fmt.Println("--- ProcessOrderPaymentHandler ---")
	fmt.Printf("Processado o pagamento de R$ %f \n", payload.TotalPrice)
	return nil
}
