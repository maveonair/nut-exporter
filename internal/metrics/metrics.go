package metrics

import (
	"github.com/maveonair/nut-exporter/internal/nut"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	probeMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nut_probe_success",
		Help: "Whether the NUT probe was successful",
	})

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

func CreateRegistryWithMetrics(state nut.UPSState) *prometheus.Registry {
	registry := CreateDefaultErrorRegistry()
	probeMetric.Set(1)

	registry.MustRegister(batteryCharge)
	batteryCharge.Set(float64(state.BatteryCharge))

	registry.MustRegister(upsLoad)
	upsLoad.Set(float64(state.UPSLoad))

	registry.MustRegister(upsRealpowerNominal)
	upsRealpowerNominal.Set(float64(state.UPSRealPowerNominal))

	registry.MustRegister(upsOnLinePower)
	if state.IsOnLinePower {
		upsOnLinePower.Set(1)
	} else {
		upsOnLinePower.Set(0)
	}

	return registry
}

func CreateDefaultErrorRegistry() *prometheus.Registry {
	errorRegistry := prometheus.NewRegistry()
	errorRegistry.MustRegister(probeMetric)
	probeMetric.Set(0)
	return errorRegistry
}
