package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wooos/alerter/internal/pkg/message"
	"github.com/wooos/alerter/internal/pkg/request"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func AlerterHandler(w http.ResponseWriter, r *http.Request) {
	requestBodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBody := request.AlertRequest{}
	if err := json.Unmarshal(requestBodyBytes, &requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	for _, alert := range requestBody.Alerts {
		message.SendMessage(alert)
	}

	fmt.Fprintf(w, "OK")
}
