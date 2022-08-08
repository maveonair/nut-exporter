package metrics

import (
	"log"
	"net/http"

	"github.com/maveonair/nut-exporter/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	batteryCharge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nut_battery_charge",
		Help: "Battery charge (percent)",
	})

	upsLoad = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nut_ups_load",
		Help: "Load on UPS (percent)",
	})

	upsRealpowerNominal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nut_ups_realpower_nominal",
		Help: "Nominal value of real power (Watts)",
	})

	upsOnLinePower = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nut_ups_on_line_power",
		Help: "Displays whether or not the ups is running on line power",
	})
)

func init() {
	prometheus.MustRegister(batteryCharge)
	prometheus.MustRegister(upsLoad)
	prometheus.MustRegister(upsRealpowerNominal)
	prometheus.MustRegister(upsOnLinePower)
}

func Serve(config config.Config) {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(config.ListeningAddr, nil))
}

func SetBatteryCharge(value int64) {
	batteryCharge.Set(float64(value))
}

func SetUpsLoad(value int64) {
	upsLoad.Set(float64(value))
}

func SetUpsRealpowerNominal(value int64) {
	upsRealpowerNominal.Set(float64(value))
}

func SetUpsOnLinePower(value int64) {
	upsOnLinePower.Set(float64(value))
}
