package main

//import "time"

type User struct {
   UserId           string    `json:"cid"`
   DisplayName      string    `json:"display_name"`
   FirstName        string    `json:"first_name"`
   Initials         string    `json:"initials"`
   LastName         string    `json:"last_name"`
   Description      string    `json:"description"`
   Department       string    `json:"department"`
   Title            string    `json:"title"`
   Office           string    `json:"office"`
   Phone            string    `json:"phone"`
   Email            string    `json:"email"`
}

