package config

import (
	"flag"
	"fmt"
	"userws/logger"
)

type Config struct {
	ServiceName       string
	Port              string
	EndpointUrl       string
	Timeout           int
	LdapBaseDn        string
	HealthCheckUser   string
	AuthTokenEndpoint string
	Debug             bool
}

var Configuration = LoadConfig()

func LoadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.Port, "port", "8080", "The service listen port")
	flag.StringVar(&c.EndpointUrl, "url", "ldap.virginia.edu:389", "The ldap hostname:port")
	flag.IntVar(&c.Timeout, "timeout", 15, "The external service timeout in seconds")
	flag.StringVar(&c.LdapBaseDn, "basedn", "o=University of Virginia,c=US", "The ldap base DN")
	flag.StringVar(&c.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
	flag.StringVar(&c.AuthTokenEndpoint, "tokenauth", "http://docker1.lib.virginia.edu:8200", "The token authentication endpoint")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	logger.Log(fmt.Sprintf("Port:                %s", c.Port))
	logger.Log(fmt.Sprintf("Endpoint:            %s", c.EndpointUrl))
	logger.Log(fmt.Sprintf("Timeout:             %d", c.Timeout))
	logger.Log(fmt.Sprintf("DN:                  %s", c.LdapBaseDn))
	logger.Log(fmt.Sprintf("Health check user:   %s", c.HealthCheckUser))
	logger.Log(fmt.Sprintf("Token auth endpoint: %s", c.AuthTokenEndpoint))
	logger.Log(fmt.Sprintf("Debug                %t", c.Debug))

	return c
}
