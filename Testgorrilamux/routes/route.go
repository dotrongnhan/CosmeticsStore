package routes

import (
	"Testgorillamux/handler"
	"Testgorillamux/middlewares"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	router.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", handler.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/cate/{category_name}", handler.GetProductByCategoryName).Methods("GET")
	router.HandleFunc("/api/products/brand/{brand_name}", handler.GetProductByBrandName).Methods("GET")
	// s := router.PathPrefix("").Subrouter()
	// s.Use(middlewares.JwtVerify)
	// router.Use(middlewares.JwtVerify)
	router.HandleFunc("/api/products", middlewares.JwtVerify(handler.CreateProduct)).Methods("POST")
	router.HandleFunc("/api/products/{id}", middlewares.JwtVerify(handler.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/api/products/{id}", middlewares.JwtVerify(handler.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/api/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/logout", handler.Logout).Methods("POST")

	router.HandleFunc("/api/brands", handler.GetBrands).Methods("GET")
	router.HandleFunc("/api/brands", handler.CreateBrand).Methods("POST")
	router.HandleFunc("/api/brands/{id}", handler.GetBrand).Methods("GET")
	router.HandleFunc("/api/brands/{id}", handler.UpdateBrand).Methods("PUT")
	router.HandleFunc("/api/brands/{id}", handler.DeleteBrand).Methods("DELETE")

	router.HandleFunc("/api/categories", handler.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories", handler.CreateCategory).Methods("POST")
	router.HandleFunc("/api/categories/{id}", handler.GetCategory).Methods("GET")
	router.HandleFunc("/api/categories/{id}", handler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/categories/{id}", handler.DeleteCategory).Methods("DELETE")
}
