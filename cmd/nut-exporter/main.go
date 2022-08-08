package main

import (
	"fmt"
	"time"

	"github.com/maveonair/nut-exporter/internal/config"
	"github.com/maveonair/nut-exporter/internal/metrics"
	nut "github.com/robbiet480/go.nut"

	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Fatal("could not initialize configuration")
	}

	go metrics.Serve(*config)

	for {
		client, connectErr := nut.Connect(config.UPSServer)
		if connectErr != nil {
			fmt.Print(connectErr)
		}

		upsList, listErr := client.GetUPSList()
		if listErr != nil {
			fmt.Print(listErr)
		}

		ups := upsList[0]

		for _, variable := range ups.Variables {
			switch variable.Name {
			case "battery.charge":
				metrics.SetBatteryCharge(variable.Value.(int64))
			case "ups.load":
				metrics.SetUpsLoad(variable.Value.(int64))
			case "ups.realpower.nominal":
				metrics.SetUpsRealpowerNominal(variable.Value.(int64))
			case "ups.status":
				value := variable.Value.(string)
				if value == "OL" {
					metrics.SetUpsOnLinePower(1)
				} else {
					metrics.SetUpsOnLinePower(0)
				}
			}
		}

		time.Sleep(config.Interval)
	}
}
