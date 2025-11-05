package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloPerson)
	http.ListenAndServe(":8080", router)
}
func HelloPerson(res http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	pss := req.FormValue("password")
	fmt.Fprintln(res, "Jana na ho", id, pss)
}
