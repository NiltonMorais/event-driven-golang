package entity

import "github.com/google/uuid"

const (
	OrderStatusPending = "pending"
)

type OrderEntity struct {
	id         string
	status     string
	items      []*OrderItemEntity
	totalPrice float64
}

func NewOrderEntity() (*OrderEntity, error) {
	return &OrderEntity{
		id:     uuid.New().String(),
		status: OrderStatusPending,
	}, nil
}

func (o *OrderEntity) AddItem(item *OrderItemEntity) {
	o.items = append(o.items, item)
	o.totalPrice += item.GetTotalPrice()
}

func (o *OrderEntity) GetItems() []*OrderItemEntity {
	return o.items
}

func (o *OrderEntity) GetTotalPrice() float64 {
	return o.totalPrice
}

func (o *OrderEntity) GetID() string {
	return o.id
}

func (o *OrderEntity) GetStatus() string {
	return o.status
}
