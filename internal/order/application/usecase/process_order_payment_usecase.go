package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/entity"
	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/queue"
)

type ProcessOrderPaymentUseCase struct {
	publisher queue.Publisher
}

func NewProcessOrderPaymentUseCase(publisher queue.Publisher) *ProcessOrderPaymentUseCase {
	return &ProcessOrderPaymentUseCase{
		publisher: publisher,
	}
}

func (h *ProcessOrderPaymentUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- ProcessOrderPaymentUseCase ---")
	order, err := entity.RestoreOrderEntity(payload.Id, payload.Status)
	if err != nil {
		return err
	}
	for _, item := range payload.Items {
		product, _ := entity.NewProductEntity(item.ProductName, item.TotalPrice/float64(item.Quantity))
		order.AddItem(entity.NewOrderItemEntity(product, item.Quantity))
	}
	paymentValue := payload.TotalPrice
	err = order.Pay(paymentValue)
	if err != nil {
		return err
	}

	fmt.Printf("Processado o pagamento de R$ %f \n", payload.TotalPrice)
	err = h.publisher.Publish(ctx, event.OrderPaidEvent{OrderId: payload.Id, PaidValue: paymentValue, PaymentDate: time.Now()})
	if err != nil {
		return err
	}
	return nil
}
