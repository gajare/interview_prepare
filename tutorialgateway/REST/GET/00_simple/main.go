package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", CheckGet).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func CheckGet(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hi amol")
}
