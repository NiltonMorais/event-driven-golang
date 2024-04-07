package event

import "github.com/NiltonMorais/event-driven-golang/internal/order/domain/entity"

type OrderCreatedEvent struct {
	Id         string
	Items      []OrderItem
	TotalPrice float64
	Status     string
}

type OrderItem struct {
	ProductName string
	Quantity    int
	TotalPrice  float64
}

func NewOrderCreatedEvent(order *entity.OrderEntity) OrderCreatedEvent {
	var items []OrderItem
	for _, item := range order.GetItems() {
		items = append(items, OrderItem{
			ProductName: item.GetProduct().GetName(),
			Quantity:    item.GetQuantity(),
			TotalPrice:  item.GetTotalPrice(),
		})
	}
	return OrderCreatedEvent{
		Id:         order.GetID(),
		TotalPrice: order.GetTotalPrice(),
		Status:     order.GetStatus(),
		Items:      items,
	}
}
