package main

import (
	"net/http"

	Admin "msebaktir.com/LCV/admin"
)

type HandleFunction func(w http.ResponseWriter, r *http.Request)
type Routes struct {
	url         string
	method      string
	handlerFunc HandleFunction
}

var routes = []Routes{
	Routes{"/Admin", "GET", Admin.AdminPage},
	Routes{"/", "POST", nil},
	Routes{"/", "PUT", nil},
	Routes{"/", "DELETE", nil},
}
