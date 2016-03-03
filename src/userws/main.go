package main

import (
   "fmt"
   "log"
   "net/http"
   "flag"
)

var config = Configuration{ }

func main( ) {

	// process command line flags and setup configuration
	flag.StringVar( &config.Port, "port", "8080", "The service listen port")
	flag.StringVar( &config.LdapUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
	flag.StringVar( &config.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
	flag.StringVar( &config.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
    flag.StringVar( &config.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")

	flag.Parse()

	log.Printf( "Port:                %s", config.Port )
	log.Printf( "LDAP endpoint:       %s", config.LdapUrl )
	log.Printf( "DN:                  %s", config.LdapBaseDn )
	log.Printf( "Health check user:   %s", config.HealthCheckUser )
    log.Printf( "Token auth endpoint: %s", config.AuthTokenEndpoint )

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Port ), router ) )
}

