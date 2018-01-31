package main

import (
	"fmt"
	"log"
	_ "expvar"
	"net/http"
	"userws/config"
	"userws/handlers"
	"userws/logger"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// setup router and serve...
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Configuration.Port), router))
}

//
// end of file
//
