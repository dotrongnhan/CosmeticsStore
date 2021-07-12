package models

type Order struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Payment   string `json:"payment"`
	Total     uint   `json:"total"`
}

type OrderItem struct {
	Id        uint `json:"id"`
	OrderId   uint `json:"order_id"`
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
