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
	flag.StringVar( &config.Port, "port", "8080", "The service listen port (default: 8080)")
	flag.StringVar( &config.LdapUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port (default: ldap.virginia.edu:389)")
	flag.StringVar( &config.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN (default: o=University of Virginia,c=US")

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Port ), router ) )
}

