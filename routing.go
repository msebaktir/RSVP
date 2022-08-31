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
	{"/", "GET", Admin.AdminPage},
	{"/admin/events", "GET", Admin.EventsPage},
}

func getAllRoutes() []Routes {
	return routes
}
