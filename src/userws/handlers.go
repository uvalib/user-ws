package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "userws/api"
    "log"
    "strings"
    "userws/ldap"
    "userws/authtoken"
    "userws/config"
)

func UserShow( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if parametersOk( userId, token ) == false {
        encodeStandardResponse(w, http.StatusBadRequest, nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        encodeStandardResponse(w, http.StatusForbidden, nil )
        return
    }

    // do the lookup
    user, err := ldap.LookupUser( config.Configuration.LdapUrl, config.Configuration.LdapBaseDn, userId )

    // lookup error?
    if err != nil {
        encodeStandardResponse( w, http.StatusInternalServerError, nil )
        return
    }

    // user not found (probably an error)?
    if user == nil {
        encodeStandardResponse( w, http.StatusNotFound, nil )
        return
    }

    // all good...
    encodeStandardResponse( w, http.StatusOK, user )
}

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

    encodeHealthCheckResponse( w, http.StatusOK, healthy, message )
}

func GetVersion( w http.ResponseWriter, r *http.Request ) {
    encodeVersionResponse( w, http.StatusOK, Version( ) )
}

func encodeStandardResponse( w http.ResponseWriter, status int, user * api.User ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: http.StatusText( status ), User: user } ); err != nil {
        log.Fatal( err )
    }
}

func encodeHealthCheckResponse( w http.ResponseWriter, status int, healthy bool, message string ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func encodeVersionResponse( w http.ResponseWriter, status int, version string ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.VersionResponse { Version: version } ); err != nil {
        log.Fatal( err )
    }
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func parametersOk( userId string, token string ) bool {
    // validate inbound parameters
    return len( strings.TrimSpace( userId ) ) != 0 &&
           len( strings.TrimSpace( token ) ) != 0

}