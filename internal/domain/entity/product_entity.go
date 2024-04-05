package entity

import "github.com/google/uuid"

type ProductEntity struct {
	id    string
	name  string
	price float64
}

func NewProductEntity(name string, price float64) (*ProductEntity, error) {
	return &ProductEntity{
		id:    uuid.New().String(),
		name:  name,
		price: price,
	}, nil
}

func (p *ProductEntity) GetID() string {
	return p.id
}

func (p *ProductEntity) GetName() string {
	return p.name
}

func (p *ProductEntity) GetPrice() float64 {
	return p.price
}
