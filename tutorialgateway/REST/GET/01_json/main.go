package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Amol"},
		{ID: 2, Name: "Rahul"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

	/*
		data, err := json.Marshal(users)        // Convert Go → JSON []byte
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)                           // Write JSON to response

	*/

}

func main() {
	router := mux.NewRouter()

	// GET route
	router.HandleFunc("/users", GetUsers).Methods("GET")

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}
