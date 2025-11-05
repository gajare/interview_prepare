package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getAll).Methods("GET")
	http.ListenAndServe(":8080", r)
}
func getAll(w http.ResponseWriter, r *http.Request) {
	Query := r.URL.Query()
	id := Query.Get("id")
	name := Query.Get("name")
	fmt.Fprintln(w, "Hi ", name, "with id ", id)
	fmt.Fprintln(w, "Hello ", name, "with id ", id)
	output := fmt.Sprintln("HI amol")
	w.Write([]byte(output))
}
