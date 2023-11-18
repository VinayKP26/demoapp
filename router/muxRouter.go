package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) Server(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/info", h.handlers.InfoController).Methods("GET")
	router.HandleFunc("/api/v1/data", h.handlers.DataController).Methods("POST")
	router.HandleFunc("/api/v1/health/check", h.handlers.HealthCheckController).Methods("GET")

	http.ListenAndServe(":"+port, router)
}
