package api

//
// User -- data associated with a user
//
type User struct {
	UserID      string   `json:"cid,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	FirstName   string   `json:"first_name,omitempty"`
	Initials    string   `json:"initials,omitempty"`
	LastName    string   `json:"last_name,omitempty"`
	Description []string `json:"description,omitempty"`
	Department  []string `json:"department,omitempty"`
	Title       []string `json:"title,omitempty"`
	Office      []string `json:"office,omitempty"`
	Phone       []string `json:"phone,omitempty"`
	Email       string   `json:"email,omitempty"`
	Private     string   `json:"private,omitempty"`
}

//
// end of file
//
