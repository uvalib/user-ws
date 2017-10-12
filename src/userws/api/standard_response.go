package api

//
// StandardResponse -- basic structure for a response
//
type StandardResponse struct {
   Status  int    `json:"status"`
   Message string `json:"message"`
   User    *User  `json:"user,omitempty"`
}

//
// end of file
//
