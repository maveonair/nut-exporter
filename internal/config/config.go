package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	envUPSserver    = "UPS_SERVER"
	envListeingAddr = "LISTENING_ADDR"
	envInterval     = "INTERVAL"
)

func init() {
	godotenv.Load()
}

type Config struct {
	ListeningAddr string
	UPSServer     string
	Interval      time.Duration
}

func NewConfig() (*Config, error) {
	var config Config

	upsServer, err := getUpsServer()
	if err != nil {
		return nil, err
	}
	config.UPSServer = upsServer

	listeningAddr, err := getListeningAddr()
	if err != nil {
		return nil, err
	}
	config.ListeningAddr = listeningAddr

	interval, err := getInterval()
	if err != nil {
		return nil, err
	}
	config.Interval = interval

	return &config, nil
}

func getUpsServer() (string, error) {
	return getEnvironmentVariable(envUPSserver)
}

func getListeningAddr() (string, error) {
	return getEnvironmentVariable(envListeingAddr)
}

func getInterval() (time.Duration, error) {
	value, err := getEnvironmentVariable(envInterval)
	if err != nil {
		return 0, err
	}

	interval, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return time.Duration(interval * int(time.Second)), nil
}

func getEnvironmentVariable(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", name)
	}

	return value, nil
}
