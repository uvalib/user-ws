package config

import (
    "flag"
    "userws/logger"
    "fmt"
)

type Config struct {
    ServiceName        string
    Port               string
    LdapUrl            string
    LdapBaseDn         string
    HealthCheckUser    string
    AuthTokenEndpoint  string
}

var Configuration = LoadConfig( )

func LoadConfig( ) Config {

    c := Config{ }

    // process command line flags and setup configuration
    flag.StringVar( &c.Port, "port", "8080", "The service listen port")
    flag.StringVar( &c.LdapUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
    flag.StringVar( &c.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
    flag.StringVar( &c.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
    flag.StringVar( &c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")

    flag.Parse()

    logger.Log( fmt.Sprintf( "Port:                %s", c.Port ) )
    logger.Log( fmt.Sprintf( "LDAP endpoint:       %s", c.LdapUrl ) )
    logger.Log( fmt.Sprintf( "DN:                  %s", c.LdapBaseDn ) )
    logger.Log( fmt.Sprintf( "Health check user:   %s", c.HealthCheckUser ) )
    logger.Log( fmt.Sprintf( "Token auth endpoint: %s", c.AuthTokenEndpoint ) )

    return c
}