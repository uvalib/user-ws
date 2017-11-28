package tests

import (
   "net/http"
   "testing"
   "userws/client"
)

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

//
// end of file
//