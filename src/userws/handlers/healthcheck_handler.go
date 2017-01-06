package handlers

import (
    "net/http"
    "userws/config"
    "userws/ldap"
)

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    healthy := true
    message := ""

    user, err := ldap.LookupUser( config.Configuration.LdapUrl, config.Configuration.LdapBaseDn, config.Configuration.HealthCheckUser )
    if err != nil {
        healthy = false
        message = err.Error( )
    } else if user == nil {
        healthy = false
    }

    encodeHealthCheckResponse( w, healthy, message )
}