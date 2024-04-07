package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/order/application/dto"
	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/entity"
	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/queue"
)

type CreateOrderUseCase struct {
	publisher queue.Publisher
}

func NewCreateOrderUseCase(publisher queue.Publisher) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		publisher: publisher,
	}
}

func (u *CreateOrderUseCase) Execute(ctx context.Context, input dto.CreateOrderDTO) error {
	fmt.Println("--- CreateOrderUseCase ---")
	order, err := entity.NewOrderEntity()

	product1, _ := entity.NewProductEntity("Product A", 10.50)
	item1 := entity.NewOrderItemEntity(product1, 1)

	product2, _ := entity.NewProductEntity("Product B", 43.19)
	item2 := entity.NewOrderItemEntity(product2, 2)

	order.AddItem(item1)
	order.AddItem(item2)

	if err != nil {
		return err
	}

	err = u.publisher.Publish(ctx, event.NewOrderCreatedEvent(order))
	if err != nil {
		return err
	}
	return nil
}
