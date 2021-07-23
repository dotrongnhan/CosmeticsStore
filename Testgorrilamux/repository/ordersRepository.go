package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetOrderItems(limit int, offset int) (orderItems []models.OrderItem, countOrders int) {
	offsetValue := limit * (offset - 1)
	query, err := database.DB.Query("SELECT * FROM order_items LIMIT ? OFFSET ?", limit, offsetValue)
	count, _ := database.DB.Query("SELECT COUNT(*) FROM order_items")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var orderItem models.OrderItem
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid, &orderItem.Ship)
		orderItem.Product = GetProduct(int(orderItem.ProductId))
		orderItem.User = GetUser(int(orderItem.UserId))
		if err != nil {
			err.Error()
		}
		orderItems = append(orderItems, orderItem)
	}

	for count.Next() {
		count.Scan(&countOrders)
	}

	return orderItems, countOrders
}

func GetOrderItem(id int) (orderItem models.OrderItem) {
	query, err := database.DB.Query("SELECT * FROM `order_items` WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid)
		if err != nil {
			err.Error()
		}
	}
	return orderItem
}

func GetOrderItemsByUserId(id int) (orderItems []models.OrderItem) {
	query, err := database.DB.Query("SELECT * FROM `order_items` WHERE user_id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var orderItem models.OrderItem
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid)
		if err != nil {
			err.Error()
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}

func CreateOrderItem(orderItem models.OrderItem) error {
	stmt, err := database.DB.Prepare("INSERT INTO order_items(user_id, product_id, quantity, is_paid) VALUES (?,?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(orderItem.UserId, orderItem.ProductId, orderItem.Quantity, orderItem.IsPaid)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateOrderItem(orderItem models.OrderItem, id int) error {
	stmt, err := database.DB.Prepare("UPDATE order_items SET product_id=?, quantity=?, is_paid=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(orderItem.ProductId, orderItem.Quantity, orderItem.IsPaid, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteOrderItem(id int, userId int) error {
	query, err := database.DB.Query("DELETE FROM order_items WHERE product_id = ? AND user_id = ? AND is_paid = 0", id, userId)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}

func DeleteOrder(id int) error {
	query, err := database.DB.Query("DELETE FROM order_items WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}

//func ChangeStatus(data )  {
//
//}
