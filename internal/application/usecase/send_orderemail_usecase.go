package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

type SendOrderEmailUseCase struct {
}

func NewSendOrderEmailUseCase() *SendOrderEmailUseCase {
	return &SendOrderEmailUseCase{}
}

func (h *SendOrderEmailUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- SendOrderEmailHandler ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
