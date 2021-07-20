package handler

import (
	"Testgorillamux/models"
	"Testgorillamux/repository"
	"Testgorillamux/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var orders []models.Order
	result := repository.GetOrderItems()
	json.NewEncoder(w).Encode(result)
}

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetOrderItem(id)
	json.NewEncoder(w).Encode(result)
}

func GetOrderItemsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetOrderItemsByUserId(id)
	json.NewEncoder(w).Encode(result)
}

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var id string
	var userId int
	cookie, _ := r.Cookie("userId")
	if cookie == nil {
		userId = 0
	} else {
		id, _ = util.ParseJwt(cookie.Value)
		userId, _ = strconv.Atoi(id)
	}
	var data map[string]uint
	json.Unmarshal(body, &data)
	orderItem := models.OrderItem{
		UserId:    uint(userId),
		ProductId: data["product_id"],
		Quantity:  data["quantity"],
		IsPaid:    int(data["is_paid"]),
	}
	errr := repository.CreateOrderItem(orderItem)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Order!")
	}
	fmt.Fprintf(w, "New OrderItem was created")
	json.NewEncoder(w).Encode(orderItem)
}

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var orderItem models.OrderItem
	json.Unmarshal(body, &orderItem)
	errr := repository.UpdateOrderItem(orderItem, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update order!")
	}
	json.NewEncoder(w).Encode(orderItem)
	fmt.Fprintf(w, "order with ID = %s was updated", params["id"])
}

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteOrderItem(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete order!")
	}
	fmt.Fprintf(w, "order item with ID = %s was deleted", params["id"])
}
