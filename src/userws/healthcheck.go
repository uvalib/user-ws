package main

type HealthCheckResult struct {
	Healthy        bool `json:"healthy"`
}

type HealthCheckResponse struct {
	CheckType      HealthCheckResult `json:"ldap"`
}

