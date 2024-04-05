package entity

type OrderItemEntity struct {
	product    *ProductEntity
	quantity   int
	totalPrice float64
}

func NewOrderItemEntity(product *ProductEntity, quantity int) *OrderItemEntity {
	return &OrderItemEntity{
		product:  product,
		quantity: quantity,
	}
}

func (o *OrderItemEntity) GetProduct() *ProductEntity {
	return o.product
}

func (o *OrderItemEntity) GetQuantity() int {
	return o.quantity
}

func (o *OrderItemEntity) GetTotalPrice() float64 {
	return o.product.GetPrice() * float64(o.quantity)
}
