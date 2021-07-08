package models

type User struct {
	Id              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	DateOfBirth     string `json:"date_of_birth"`
	Gender          string `json:"gender"`
	Avatar          string `json:"avatar"`
	RoleId          uint   `json:"role_id"`
}
