package userws

//import "time"

type Response struct {
   Status        int     `json:"status"`
   Message       string  `json:"message"`
   User          User    `json:"user,omitempty"`
}

