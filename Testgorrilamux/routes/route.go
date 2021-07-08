package routes

import (
	"Testgorillamux/handler"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	router.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/api/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", handler.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", handler.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/api/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/logout", handler.Logout).Methods("POST")

	// router.HandleFunc("/pictures", getPictures).Methods("GET")
	// router.HandleFunc("/pictures", createPicture).Methods("POST")
	// router.HandleFunc("/pictures/{id}", getPicture).Methods("GET")
	// router.HandleFunc("/pictures/{id}", updatePicture).Methods("PUT")
	// router.HandleFunc("/pictures/{id}", deletePicture).Methods("DELETE")

	// router.HandleFunc("/categories", getCategories).Methods("GET")
	// router.HandleFunc("/categories", createCategory).Methods("POST")
	// router.HandleFunc("/categories/{id}", getCategory).Methods("GET")
	// router.HandleFunc("/categories/{id}", updateCategory).Methods("PUT")
	// router.HandleFunc("/categories/{id}", deleteCategory).Methods("DELETE")
}
