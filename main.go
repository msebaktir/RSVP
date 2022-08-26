package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	for _, element := range getAllRoutes() {
		r.HandleFunc(element.url, element.handlerFunc).Methods(element.method)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Listening on port 8080 ... => http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
