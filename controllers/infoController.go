package controllers

import "net/http"

func (c *Controllers) InfoController(w http.ResponseWriter, r *http.Request) {
	body := "This is demoapp. This api is providing information to you"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
