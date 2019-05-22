package service

import "net/http"

// Route is a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array slice of Route
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",            // Name
		http.MethodGet,          // HTTP method
		"/accounts/{accountId}", // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
