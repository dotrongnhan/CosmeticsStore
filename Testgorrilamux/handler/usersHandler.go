package handler

import (
	"Testgorillamux/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := repository.GetUsers()
	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetUser(id)
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteUser(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete User!")
	}
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}
