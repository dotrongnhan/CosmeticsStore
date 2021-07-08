package handler

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []models.Product

	methods := r.URL.Query()
	method, _ := strconv.Atoi(methods.Get("limit"))

	result, err := database.DB.Query("SELECT P.id, P.title, P.description, P.image, P.price, P.category_id, P.is_sale, P.is_hot, P.sale_price FROM products P LIMIT ?", method)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var product models.Product
		err := result.Scan(&product.Id, &product.Title, &product.Description, &product.Image, &product.Price, &product.CategoryId, &product.IsSale, &product.IsHot, &product.SalePrice)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := database.DB.Prepare("INSERT INTO products(title, description, image, price, category_id, is_sale, is_hot, sale_price) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var product models.Product
	json.Unmarshal(body, &product)
	_, err = stmt.Exec(product.Title, product.Description, product.Image, product.Price, product.CategoryId, product.IsSale, product.IsHot, product.SalePrice)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(product)
	fmt.Fprintf(w, "New product was created")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := database.DB.Query("SELECT id, title, description, image, price, category_id, is_sale, is_hot, sale_price FROM products WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var product models.Product

	for result.Next() {
		err := result.Scan(&product.Id, &product.Title, &product.Description, &product.Image, &product.Price, &product.CategoryId, &product.IsSale, &product.IsHot, &product.SalePrice)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := database.DB.Prepare("UPDATE products SET title=?,description=?,image=?,price=?,category_id=?,is_sale=?,is_hot=?,sale_price=? WHERE id = ?;")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var product models.Product
	json.Unmarshal(body, &product)
	_, err = stmt.Exec(product.Title, product.Description, product.Image, product.Price, product.CategoryId, product.IsSale, product.IsHot, product.SalePrice, params["id"])
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(product)
	fmt.Fprintf(w, "Product with ID = %s was updated", params["id"])
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := database.DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Product with ID = %s was deleted", params["id"])
}
