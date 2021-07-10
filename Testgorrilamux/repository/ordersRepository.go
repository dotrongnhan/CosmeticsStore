package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetOrders() (orders []models.Order) {
	query, err := database.DB.Query("SELECT * FROM orders")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	for query.Next() {
		var order models.Order
		err := query.Scan(&order.Id, &order.UserId, &order.Total)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders
}

func GetOrder(id int) (order models.Order) {
	query, err := database.DB.Query("SELECT * FROM orders WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&order.Id, &order.UserId, &order.Total)
		if err != nil {
			panic(err.Error())
		}
	}
	return order
}

func CreateOrder(order models.Order) error {
	stmt, err := database.DB.Prepare("INSERT INTO Orders(user_id, total) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(order.UserId, order.Total)
	if err != nil {
		panic(err.Error())
	}
	return err
}

func UpdateOrder(order models.Order, id int) error {
	stmt, err := database.DB.Prepare("UPDATE orders SET user_id=?, total=? WHERE id = ?;")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(order.UserId, order.Total, id)
	if err != nil {
		panic(err.Error())
	}
	return err
}

func DeleteOrder(id int) error {
	query, err := database.DB.Query("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	return err
}
