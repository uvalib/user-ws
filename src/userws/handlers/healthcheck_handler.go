package handlers

import (
	"net/http"
	"userws/config"
	"userws/ldap"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthy := true
	message := ""

	user, err := ldap.LookupUser(config.Configuration.EndpointURL,
		config.Configuration.ServiceTimeout,
		config.Configuration.LdapBaseDn,
		config.Configuration.HealthCheckUser)
	if err != nil {
		healthy = false
		message = err.Error()
	} else if user == nil {
		healthy = false
	}

	encodeHealthCheckResponse(w, healthy, message)
}

//
// end of file
//
