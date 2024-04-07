package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type ProcessOrderPaymentUseCase struct {
}

func NewProcessOrderPaymentUseCase() *ProcessOrderPaymentUseCase {
	return &ProcessOrderPaymentUseCase{}
}

func (h *ProcessOrderPaymentUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- ProcessOrderPaymentUseCase ---")
	fmt.Printf("Processado o pagamento de R$ %f \n", payload.TotalPrice)
	return nil
}
