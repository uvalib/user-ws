package main

import (
   "net/http"
   "github.com/gorilla/mux"
        "userws/handlers"
)

type Route struct {
   Name        string
   Method      string
   Pattern     string
   HandlerFunc http.HandlerFunc
}

type Routes [] Route

var routes = Routes{

    Route{
       "UserLookup",
       "GET",
       "/user/{userId}",
       handlers.UserLookup,
    },

    Route{
       "HealthCheck",
       "GET",
       "/healthcheck",
       handlers.HealthCheck,
    },

    Route{
        "VersionInfo",
        "GET",
        "/version",
        handlers.VersionInfo,
    },

    Route{
        "RuntimeInfo",
        "GET",
        "/runtime",
        handlers.RuntimeInfo,
    },
}

func NewRouter( ) *mux.Router {

   router := mux.NewRouter().StrictSlash( true )
   for _, route := range routes {

      var handler http.Handler

      handler = route.HandlerFunc
      handler = HandlerLogger( handler, route.Name )

      router.
         Methods( route.Method ).
         Path( route.Pattern ).
         Name( route.Name ).
         Handler( handler )
   }

   return router
}
