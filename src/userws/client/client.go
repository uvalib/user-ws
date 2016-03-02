package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
)

func HealthCheck( endpoint string ) int {

    url := fmt.Sprintf( "%s/healthcheck", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
    SetDebug( false ).
    Get( url  ).
    Timeout( time.Duration( 5 ) * time.Second ).
    End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}

func UserDetails( endpoint string, username string, token string ) int {

    url := fmt.Sprintf( "%s/user/%s?auth=%s", endpoint, username, token )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}