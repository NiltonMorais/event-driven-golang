package handler

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type StockMovementHandler struct {
}

func NewStockMovementHandler() *StockMovementHandler {
	return &StockMovementHandler{}
}

func (h *StockMovementHandler) Execute(ctx context.Context, e event.DomainEvent) error {
	payload, err := event.GetPayload[event.OrderCreatedEvent](e)
	if err != nil {
		return err
	}
	fmt.Println("--- StockMovimentHandler ---")
	for _, item := range payload.Items {
		fmt.Printf("Retirando do stock %d itens do produto: %s\n", item.Quantity, item.ProductName)
	}
	return nil
}
