package models

type Data struct {
	UserId 				int 					`json:"user_id"`
	Products 			[]Product 		`json:"products"`
	AddressShipping 	AddressShipping 	`json:"address_shipping"`
	Orders 				[]OrderItem		`json:"orders"`
	Total				int 					`json:"total"`
}
