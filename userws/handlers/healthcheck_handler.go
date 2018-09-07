package handlers

import (
	"fmt"
	"github.com/uvalib/user-ws/userws/config"
	"github.com/uvalib/user-ws/userws/ldap"
	"github.com/uvalib/user-ws/userws/logger"
	"net/http"
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
		logger.Log(fmt.Sprintf("ERROR: LDAP lookup reports '%s'", message))
	} else if user == nil {
		healthy = false
		logger.Log(fmt.Sprintf("ERROR: LDAP lookup cannot find '%s'", config.Configuration.HealthCheckUser))
	}

	encodeHealthCheckResponse(w, healthy, message)
}

//
// end of file
//
