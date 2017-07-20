package api

type StandardResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	User    *User  `json:"user,omitempty"`
}
