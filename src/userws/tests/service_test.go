package tests

import (
   "io/ioutil"
   "log"
   "strings"
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