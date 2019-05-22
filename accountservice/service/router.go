package service

import (
	"github.com/gorilla/mux"
)

// NewRoute generates routes with Gorilla Router
func NewRoute() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
