package tests

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type testConfig struct {
	TestUser string
	Endpoint string
	Secret   string
}

var cfg = loadConfig()

var goodUser = cfg.TestUser
var badUser = "xxyyzz"

// var goodToken = cfg.Token
// var badToken = "badness"
var empty = " "

func emptyFields(fields []string) bool {

	for _, field := range fields {
		if emptyField(field) == true {
			return true
		}
	}
	return false
}

func emptyField(field string) bool {
	return len(strings.TrimSpace(field)) == 0
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("testuser [%s]\n", c.TestUser)
	log.Printf("endpoint [%s]\n", c.Endpoint)
	log.Printf("secret   [%s]\n", c.Secret)

	return c
}

func badToken(secret string) string {

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(-5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func goodToken(secret string) string {

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

//
// end of file
//
