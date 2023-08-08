package server

import (
	"net/http"

	"github.com/maveonair/nut-exporter/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

func (s *server) probeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("target")

		var registry *prometheus.Registry

		state, err := s.nutClient.GetUPSState(target)
		if err != nil {
			log.WithFields(log.Fields{
				"target": target,
			}).WithError(err).Error("failed to get UPS state")

			registry = metrics.CreateDefaultErrorRegistry()
		} else {
			registry = metrics.CreateRegistryWithMetrics(state)
		}

		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		handler.ServeHTTP(w, r)
	}
}
