package controllers

import "net/http"

func (c *Controllers) HealthCheckController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
