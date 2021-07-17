package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetUsers() (customers []models.Customer) {
	query, err := database.DB.Query("SELECT id, full_name, email, phone, address, date_of_birth, gender, avatar, role_id FROM users WHERE role_id = 2")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var customer models.Customer
		err := query.Scan(&customer.Id, &customer.FullName, &customer.Email, &customer.Phone, &customer.Address, &customer.DateOfBirth, &customer.Gender, &customer.Avatar, &customer.RoleId)
		if err != nil {
			err.Error()
		}
		customers = append(customers, customer)
	}
	return customers
}

func GetUser(id int) (customer models.Customer) {
	query, err := database.DB.Query("SELECT id, full_name, email, phone, address, date_of_birth, gender, avatar, role_id FROM users WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&customer.Id, &customer.FullName, &customer.Email, &customer.Phone, &customer.Address, &customer.DateOfBirth, &customer.Gender, &customer.Avatar, &customer.RoleId)
		if err != nil {
			err.Error()
		}
	}
	return customer
}

func DeleteUser(id int) error {
	query, err := database.DB.Query("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
