package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}/{name}", URLCheck).Methods("GET")
	http.ListenAndServe(":8080", r)
}
func URLCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	name := vars["name"]
	fmt.Fprintln(w, "Hi id :", id, "name :", name)
}
