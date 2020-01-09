package handlers

import (
	"net/http"
)

// healthCheck is used to check the service health
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sample-fb is up and running..!!!"))
}
