package config

import (
    "flag"
    "log"
)

type Config struct {
    Port               string
    LdapUrl            string
    LdapBaseDn         string
    HealthCheckUser    string
    AuthTokenEndpoint  string
}

var Configuration = LoadConfig( )

func LoadConfig( ) Config {

    c := Config{}

    // process command line flags and setup configuration
    flag.StringVar( &c.Port, "port", "8080", "The service listen port")
    flag.StringVar( &c.LdapUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
    flag.StringVar( &c.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
    flag.StringVar( &c.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
    flag.StringVar( &c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")

    flag.Parse()

    log.Printf( "Port:                %s", c.Port )
    log.Printf( "LDAP endpoint:       %s", c.LdapUrl )
    log.Printf( "DN:                  %s", c.LdapBaseDn )
    log.Printf( "Health check user:   %s", c.HealthCheckUser )
    log.Printf( "Token auth endpoint: %s", c.AuthTokenEndpoint )

    return c
}