package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "strings"
   "testing"
   "userws/client"
   "gopkg.in/yaml.v2"
)

type testConfig struct {
   TestUser string
   Endpoint string
   Token    string
}

var cfg = loadConfig()

var goodUser = cfg.TestUser
var badUser = "xxyyzz"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

func TestHealthCheck(t *testing.T) {
   expected := http.StatusOK
   status := client.HealthCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestVersionCheck(t *testing.T) {
   expected := http.StatusOK
   status, version := client.VersionCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if len(version) == 0 {
      t.Fatalf("Expected non-zero length version string\n")
   }
}

func TestRuntimeCheck(t *testing.T) {
   expected := http.StatusOK
   status, runtime := client.RuntimeCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if runtime == nil {
      t.Fatalf("Expected non-nil runtime info\n")
   }

   if len( runtime.Version ) == 0 ||
      runtime.AllocatedMemory == 0 ||
      runtime.CPUCount == 0 ||
      runtime.GoRoutineCount == 0 ||
      runtime.ObjectCount == 0 {
      t.Fatalf("Expected non-zero value in runtime info but one is zero\n")
   }
}

func TestUserDetailsHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, user := client.UserDetails(cfg.Endpoint, goodUser, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if user == nil {
      t.Fatalf("Expected to find user %v and did not\n", goodUser)
   }

   if emptyField(user.UserID) ||
      emptyField(user.DisplayName) ||
      emptyField(user.FirstName) ||
      emptyField(user.Initials) ||
      emptyField(user.LastName) ||
      emptyField(user.Description) ||
      emptyField(user.Department) ||
      emptyField(user.Title) ||
      emptyField(user.Office) ||
      emptyField(user.Phone) ||
      emptyField(user.Email) {
      t.Fatalf("Expected non-empty field but one is empty\n")
   }
}

func TestUserDetailsEmptyUser(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.UserDetails(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestUserDetailsBadUser(t *testing.T) {
   expected := http.StatusNotFound
   status, _ := client.UserDetails(cfg.Endpoint, badUser, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestUserDetailsEmptyToken(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.UserDetails(cfg.Endpoint, goodUser, empty)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestUserDetailsBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.UserDetails(cfg.Endpoint, goodUser, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
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
   log.Printf("token    [%s]\n", c.Token)

   return c
}

//
// end of file
//