package config

import (
	"flag"
	"fmt"
	"userws/logger"
)

//
// Config -- our configuration structure
type Config struct {
	ServiceName       string
	ServicePort       string
	EndpointURL       string
	Timeout           int
	LdapBaseDn        string
	HealthCheckUser   string
	AuthTokenEndpoint string
	Debug             bool
}

//
// Configuration -- our configuration instance
//
var Configuration = loadConfig()

func loadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.EndpointURL, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
	flag.IntVar(&c.Timeout, "timeout", 15, "The external service timeout in seconds")
	flag.StringVar(&c.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
	flag.StringVar(&c.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
	flag.StringVar(&c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("ServicePort:         %s", c.ServicePort))
	logger.Log(fmt.Sprintf("Endpoint:            %s", c.EndpointURL))
	logger.Log(fmt.Sprintf("Timeout:             %d", c.Timeout))
	logger.Log(fmt.Sprintf("DN:                  %s", c.LdapBaseDn))
	logger.Log(fmt.Sprintf("Health check user:   %s", c.HealthCheckUser))
	logger.Log(fmt.Sprintf("Token auth endpoint: %s", c.AuthTokenEndpoint))
	logger.Log(fmt.Sprintf("Debug                %t", c.Debug))

	return c
}

//
// end of file
//
