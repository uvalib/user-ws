package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"userws/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
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

		var handler http.Handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)
		handler = prometheus.InstrumentHandler( route.Name, handler )

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// add the route for the prometheus metrics
	router.Handle("/metrics", HandlerLogger( promhttp.Handler( ), "promhttp.Handler" ) )
	return router
}

//
// end of file
//
