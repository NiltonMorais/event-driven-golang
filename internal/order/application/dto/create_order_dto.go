package dto

type CreateOrderDTO struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id  string `json:"id"`
	Qtd string `json:"qtd"`
}
