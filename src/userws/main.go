package main

import (
    "fmt"
    "log"
    "net/http"
    "userws/config"
    "userws/logger"
    "userws/handlers"
)

func main( ) {

    logger.Log( fmt.Sprintf( "===> version: '%s' <===", handlers.Version( ) ) )

    // setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.Port ), router ) )
}

