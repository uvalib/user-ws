package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"userws/handlers"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routeSlice []route

var routes = routeSlice{

	route{
		"UserLookup",
		"GET",
		"/user/{userId}",
		handlers.UserLookup,
	},

	route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.HealthCheck,
	},

	route{
		"VersionInfo",
		"GET",
		"/version",
		handlers.VersionInfo,
	},

	route{
		"RuntimeInfo",
		"GET",
		"/runtime",
		handlers.RuntimeInfo,
	},
}

//
// NewRouter -- build and return the router
//
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// add the route for the expvars eendpoint
	router.Handle("/debug/vars", http.DefaultServeMux )

	return router
}

//
// end of file
//
