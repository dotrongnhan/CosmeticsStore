package handler

import (
	"Testgorillamux/models"
	"Testgorillamux/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type resProduct struct {
	Products	[]models.Product
	Count		int
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var products []models.Product
	limit := 16
	params := r.URL.Query()
	sortType := params.Get("sortType")
	offset, _ := strconv.Atoi(params.Get("offset"))
	category := params.Get("category")
	result, count := repository.GetProducts(sortType, category, limit, offset)
	json.NewEncoder(w).Encode(resProduct{result, count})
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var data map[string]string
	json.Unmarshal(body, &data)
	result := repository.SearchProducts(data["key_word"])
	json.NewEncoder(w).Encode(result)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var product models.Product
	json.Unmarshal(body, &product)
	error := repository.CreateProduct(product)
	if error != nil {
		fmt.Fprintln(w, "Cannot create product!")
	}
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetProduct(id)
	json.NewEncoder(w).Encode(result)
}

func GetProductByCategoryName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["category_name"]
	result := repository.GetProductByCategoryName(name)
	json.NewEncoder(w).Encode(result)
}

func GetProductByBrandName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["brand_name"]
	result := repository.GetProductByCategoryName(name)
	json.NewEncoder(w).Encode(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var product models.Product
	json.Unmarshal(body, &product)
	errr := repository.UpdateProduct(product, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update product!")
	}
	json.NewEncoder(w).Encode(product)
	fmt.Fprintf(w, "Product with ID = %s was updated", params["id"])
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteProduct(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete product!")
	}
	fmt.Fprintf(w, "Product with ID = %s was deleted", params["id"])
}
