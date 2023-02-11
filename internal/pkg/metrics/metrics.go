package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	AlerterSendMessageTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "alerter_send_message_total",
			Help: "Total for send message",
		},
		[]string{"channel", "status"},
	)
)

func NewMetricsHandler() http.Handler {
	prometheus.MustRegister(
		AlerterSendMessageTotal,
	)

	return promhttp.Handler()
}
