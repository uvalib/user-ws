package main

import (
   "fmt"
   "log"
   "net/http"
   "flag"
)

var config = Configuration{ }

func main( ) {

	// process command line flags and serup configuration
	flag.StringVar( &config.Port, "port", "8080", "The service listen port")
	flag.StringVar( &config.LdapUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
	flag.StringVar( &config.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
	flag.Parse()

	log.Printf( "Port: %s", config.Port )
	log.Printf( "URL:  %s", config.LdapUrl )
	log.Printf( "DN:   %s", config.LdapBaseDn )

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Port ), router ) )
}

