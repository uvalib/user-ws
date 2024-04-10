package api

// HealthCheckResponse -- response to the health check query
type HealthCheckResponse struct {
	CheckType HealthCheckResult `json:"ldap"`
}

//
// end of file
//
