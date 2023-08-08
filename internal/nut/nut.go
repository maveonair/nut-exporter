package nut

import (
	nut "github.com/robbiet480/go.nut"
)

type UPSState struct {
	BatteryCharge       int64
	IsOnLinePower       bool
	UPSLoad             int64
	UPSRealPowerNominal int64
}

type Client interface {
	GetUPSState(upsServer string) (UPSState, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (client) GetUPSState(upsServer string) (UPSState, error) {
	client, err := nut.Connect(upsServer)
	if err != nil {
		return UPSState{}, err
	}

	upsList, err := client.GetUPSList()
	if err != nil {
		return UPSState{}, err
	}

	ups := upsList[0]

	state := UPSState{}

	for _, variable := range ups.Variables {
		switch variable.Name {
		case "battery.charge":
			state.BatteryCharge = variable.Value.(int64)
		case "ups.load":
			state.UPSLoad = variable.Value.(int64)
		case "ups.realpower.nominal":
			state.UPSRealPowerNominal = variable.Value.(int64)
		case "ups.status":
			value := variable.Value.(string)
			if value == "OL" {
				state.IsOnLinePower = true
			} else {
				state.IsOnLinePower = false
			}
		}
	}

	return state, nil
}
