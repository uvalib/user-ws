package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "testing"
    "userws/client"
    "gopkg.in/yaml.v2"
    "net/http"
)

type TestConfig struct {
    Endpoint  string
    Token     string
}

var cfg = loadConfig( )

var goodUser = "dpg3k"
var badUser = "xxyyzz"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

func TestHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status := client.UserDetails( cfg.Endpoint, goodUser, goodToken )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func TestEmptyUser( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.UserDetails( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func TestBadUser( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.UserDetails( cfg.Endpoint, badUser, goodToken )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func TestEmptyToken( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.UserDetails( cfg.Endpoint, goodUser, empty )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func TestBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status := client.UserDetails( cfg.Endpoint, goodUser, badToken )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func TestHealthCheck( t *testing.T ) {
    expected := http.StatusOK
    status := client.HealthCheck( cfg.Endpoint )
    if status != expected {
        t.Errorf( "Expected %v, got %v\n", expected, status )
    }
}

func loadConfig( ) TestConfig {

    data, err := ioutil.ReadFile( "service_test.yml" )
    if err != nil {
        log.Fatal( err )
    }

    var c TestConfig
    if err := yaml.Unmarshal( data, &c ); err != nil {
        log.Fatal( err )
    }

    fmt.Printf( "endpoint [%s]\n", c.Endpoint )
    fmt.Printf( "token    [%s]\n", c.Token )

    return c
}