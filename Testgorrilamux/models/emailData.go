package models

type EmailData struct {
	Orders 				[]OrderItem		`json:"orders"`
	Total				int 					`json:"total"`
}
