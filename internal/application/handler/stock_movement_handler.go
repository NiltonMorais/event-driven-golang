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

func (h *StockMovementHandler) Execute(ctx context.Context, e event.DomainEvent) {
	payload := event.GetPayload[event.OrderCreatedEvent](e)
	fmt.Println("--- StockMovimentHandler ---")
	for _, item := range payload.Order.GetItems() {
		fmt.Printf("Retirando do stock %d itens do produto: %s\n", item.GetQuantity(), item.GetProduct().GetName())
	}
}
