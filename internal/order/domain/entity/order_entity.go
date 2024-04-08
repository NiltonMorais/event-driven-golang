package entity

import (
	"errors"

	"github.com/google/uuid"
)

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type OrderEntity struct {
	id         string
	status     string
	items      []*OrderItemEntity
	totalPrice float64
	paidValue  float64
}

func NewOrderEntity() (*OrderEntity, error) {
	return &OrderEntity{
		id:     uuid.New().String(),
		status: OrderStatusPending,
	}, nil
}

func RestoreOrderEntity(id, status string) (*OrderEntity, error) {
	return &OrderEntity{
		id:     id,
		status: status,
	}, nil
}

func (o *OrderEntity) AddItem(item *OrderItemEntity) {
	o.items = append(o.items, item)
	o.totalPrice += item.GetTotalPrice()
}

func (o *OrderEntity) Pay(value float64) error {
	if value < o.totalPrice {
		return errors.New("value is less than the total price")
	}
	o.paidValue = value
	o.status = OrderStatusPaid
	return nil
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
