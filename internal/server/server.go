package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wooos/alerter/internal/config"
)

func RunServer(conf *config.Config) {
	server := mux.NewRouter()
	server.HandleFunc("/v1/send", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello")
	}).Methods("POST")

	srv := &http.Server{
		Handler: server,
		Addr:    conf.ListenAddr(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
