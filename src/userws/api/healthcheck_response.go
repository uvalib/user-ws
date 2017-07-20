package api

type HealthCheckResponse struct {
	CheckType HealthCheckResult `json:"ldap"`
}
