package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wooos/alerter/internal/config"
)

func RunServer() {
	server := mux.NewRouter()

	server.HandleFunc("/healthz", HealthHandler).Methods("GET")

	api := server.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/alerter", AlerterHandler).Methods("POST")

	srv := &http.Server{
		Handler: server,
		Addr:    config.Conf.ListenAddr(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting ...")
	srv.ListenAndServe()
}
