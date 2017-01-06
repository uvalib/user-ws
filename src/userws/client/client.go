package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "userws/api"
    "encoding/json"
    "io"
    "io/ioutil"
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

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    return resp.StatusCode
}

func VersionCheck( endpoint string ) ( int, string ) {

    url := fmt.Sprintf( "%s/version", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, ""
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.VersionResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, ""
    }

    return resp.StatusCode, r.Version
}

func RuntimeCheck( endpoint string ) ( int, * api.RuntimeResponse ) {

    url := fmt.Sprintf( "%s/runtime", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
            SetDebug( false ).
            Get( url  ).
            Timeout( time.Duration( 5 ) * time.Second ).
            End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.RuntimeResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, &r
}

func UserDetails( endpoint string, username string, token string ) ( int, * api.User ) {

    url := fmt.Sprintf( "%s/user/%s?auth=%s", endpoint, username, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError, nil
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.User
}