package config

import (
	"flag"
	"fmt"
	"github.com/uvalib/user-ws/userws/logger"
	"strings"
)

// Config -- our configuration structure
type Config struct {
	ServiceName       string
	ServicePort       string
	ServiceTimeout    int
	LdapEndpoint      string
	LdapBindAccount   string
	LdapBindPassword  string
	LdapBaseDn        string
	LdapUseTls        bool
	LdapSkipTlsVerify bool
	HealthCheckUser   string
	SharedSecret      string
	Debug             bool
}

// Configuration -- our configuration instance
var Configuration = loadConfig()

func loadConfig() Config {

	c := Config{}

	// process command line flags and setup configuration
	flag.StringVar(&c.ServicePort, "port", "8080", "The service listen port")
	flag.StringVar(&c.LdapEndpoint, "ldapendpoint", "ldap.virginia.edu:389", "The ldap hostname:port")
	flag.IntVar(&c.ServiceTimeout, "timeout", 15, "The external service timeout in seconds")
	flag.StringVar(&c.LdapBindAccount, "ldapbindacct", "", "The ldap bind account name")
	flag.StringVar(&c.LdapBindPassword, "ldapbindpwd", "", "The ldap bind password")
	flag.StringVar(&c.LdapBaseDn, "ldapbasedn", "o=University of Virginia,c=US", "The ldap base DN")
	flag.BoolVar(&c.LdapUseTls, "ldaptls", false, "Use ldap TLS")
	flag.BoolVar(&c.LdapSkipTlsVerify, "tlsskipverify", false, "Skip TLS certificate verification")
	flag.StringVar(&c.HealthCheckUser, "hcuser", "dpg3k", "The search name used for the health check")
	flag.StringVar(&c.SharedSecret, "secret", "", "The JWT shared secret")
	flag.BoolVar(&c.Debug, "debug", false, "Enable debugging")

	flag.Parse()

	// handle special cases here
	c.LdapBindAccount = strings.Replace(c.LdapBindAccount, "%20", " ", -1)
	c.LdapBindAccount = strings.Replace(c.LdapBindAccount, "%3D", "=", -1)
	c.LdapBaseDn = strings.Replace(c.LdapBaseDn, "%20", " ", -1)
	c.LdapBaseDn = strings.Replace(c.LdapBaseDn, "%3D", "=", -1)

	logger.Log(fmt.Sprintf("INFO: ServicePort:         %s", c.ServicePort))
	logger.Log(fmt.Sprintf("INFO: ServiceTimeout:      %d", c.ServiceTimeout))
	logger.Log(fmt.Sprintf("INFO: LdapEndpoint:        %s", c.LdapEndpoint))
	logger.Log(fmt.Sprintf("INFO: LdapBindAccount:     %s", c.LdapBindAccount))
	logger.Log(fmt.Sprintf("INFO: LdapBindPassword:    %s", strings.Repeat("*", len(c.LdapBindPassword))))
	logger.Log(fmt.Sprintf("INFO: LdapBaseDn:          %s", c.LdapBaseDn))
	logger.Log(fmt.Sprintf("INFO: LdapUseTls:          %t", c.LdapUseTls))
	logger.Log(fmt.Sprintf("INFO: LdapSkipTlsVerify:   %t", c.LdapSkipTlsVerify))
	logger.Log(fmt.Sprintf("INFO: Health check user:   %s", c.HealthCheckUser))
	logger.Log(fmt.Sprintf("INFO: SharedSecret:        %s", strings.Repeat("*", len(c.SharedSecret))))
	logger.Log(fmt.Sprintf("INFO: Debug:               %t", c.Debug))

	return c
}

//
// end of file
//
