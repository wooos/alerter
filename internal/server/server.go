package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/pkg/metrics"
)

func RunServer() {
	server := mux.NewRouter()
	server.Use()

	server.HandleFunc("/healthz", HealthHandler).Methods("GET")
	server.Handle("/metrics", metrics.NewMetricsHandler())

	api := server.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/alerter", AlerterHandler).Methods("POST")

	log.Println("Server starting ...")
	http.ListenAndServe(config.Conf.ListenAddr(), handlers.LoggingHandler(os.Stdout, server))
}
