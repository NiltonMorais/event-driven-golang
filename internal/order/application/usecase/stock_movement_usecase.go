package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
)

type StockMovementUseCase struct {
}

func NewStockMovementUseCase() *StockMovementUseCase {
	return &StockMovementUseCase{}
}

func (h *StockMovementUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- StockMovimentHandler ---")
	for _, item := range payload.Items {
		fmt.Printf("Retirando do stock %d itens do produto: %s\n", item.Quantity, item.ProductName)
	}
	return nil
}
