package models

type Order struct {
	Id     uint `json:"id"`
	UserId uint `json:"user"`
	Total  uint `json:"total"`
}

type OderItem struct {
	Id        uint `json:"id"`
	OrderId   uint `json:"order_id"`
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
