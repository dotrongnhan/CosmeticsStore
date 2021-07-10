package main

import (
	"Testgorillamux/database"
	"Testgorillamux/routes"
	"fmt"
	"github.com/rs/cors"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect() //Khởi tạo biến conection
	if err != nil {               //Catch error trong quá trình thực thi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	router := mux.NewRouter()


	//router.Use(loggingMiddleware)
	routes.Setup(router)
	handler := cors.Default().Handler(router)

	fmt.Println("Server running at localhost:8000")

	http.ListenAndServe(":8000", handler)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
