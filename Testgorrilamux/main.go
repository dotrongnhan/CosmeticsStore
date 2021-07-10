package main

import (
	"Testgorillamux/database"
	"fmt"
	"log"

	"Testgorillamux/routes"

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

	// cors := handlers.CORS(
	// 	handlers.AllowedHeaders([]string{"Content-Type", "JWT", "Set-Cookie"}),
	// 	handlers.AllowedOrigins([]string{"*"}),
	// 	handlers.AllowCredentials(),
	// )
	// router.Use(cors)

	routes.Setup(router)
	fmt.Println("Server running at localhost:8000")

	http.ListenAndServe(":8000", router)

}
