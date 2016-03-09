package main

import (
    "fmt"
    "log"
    "net/http"
    "userws/config"
)

func main( ) {

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.Port ), router ) )
}

