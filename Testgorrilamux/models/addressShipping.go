package models

type AddressShipping struct {
	Id 		uint 	`json:"id"`
	UserId 	int 	`json:"user_id"`
	Address string 	`json:"address"`
	Email 	string 	`json:"email"`
	Phone 	string 	`json:"phone"`
	FullName 	string 	`json:"full_name"`
}