package authtoken

import (
    "fmt"
    "time"
    "net/http"
    "github.com/parnurzeal/gorequest"
    "userws/logger"
    "io"
    "io/ioutil"
)

func Validate( endpoint string, timeout int, token string ) bool {

    url := fmt.Sprintf( "%s/authorize/%s/%s/%s", endpoint, "userservice", "userlookup", token )
    //log.Printf( "%s\n", url )

    start := time.Now( )
    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( timeout ) * time.Second ).
       End( )
    duration := time.Since( start )

    if errs != nil {
        logger.Log( fmt.Sprintf( "ERROR: token auth (%s) returns %s in %s\n", url, errs, duration ) )
        return false
    }

    defer io.Copy( ioutil.Discard, resp.Body )
    defer resp.Body.Close( )

    logger.Log( fmt.Sprintf( "Token auth (%s) returns http %d in %s\n", url, resp.StatusCode, duration ) )
    return resp.StatusCode == http.StatusOK
}
