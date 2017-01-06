package handlers

import (
    "encoding/json"
    "net/http"
    "userws/api"
    "log"
    "strings"
)

func encodeStandardResponse( w http.ResponseWriter, status int, user * api.User ) {
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: http.StatusText( status ), User: user } ); err != nil {
        log.Fatal( err )
    }
}

func encodeHealthCheckResponse( w http.ResponseWriter, healthy bool, message string ) {
    status := http.StatusOK
    if healthy == false {
        status = http.StatusInternalServerError
    }
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func encodeVersionResponse( w http.ResponseWriter, status int, version string ) {
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.VersionResponse { Version: version } ); err != nil {
        log.Fatal( err )
    }
}

func jsonAttributes( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func parametersOk( userId string, token string ) bool {
    // validate inbound parameters
    return len( strings.TrimSpace( userId ) ) != 0 &&
           len( strings.TrimSpace( token ) ) != 0

}