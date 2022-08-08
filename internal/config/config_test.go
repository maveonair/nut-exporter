package config

import (
	"os"
	"testing"
	"time"
)

func setup() {
	os.Clearenv()
}

func teardown() {
	os.Clearenv()
}

func Test_GetUpsServer(t *testing.T) {
	setup()

	_, err := getUpsServer()
	if err == nil {
		t.Error("getUpsServer should return an error")
	}

	expectedValue := "127.0.0.1"

	os.Setenv(envUPSserver, expectedValue)

	value, err := getUpsServer()
	if err != nil {
		t.Error("getUpsServer should not return an error")
	}

	if value != expectedValue {
		t.Errorf("expected %s, actual %s", expectedValue, value)
	}

	teardown()
}

func Test_GetListeningAddr(t *testing.T) {
	setup()

	_, err := getListeningAddr()
	if err == nil {
		t.Error("getListeningAddr should return an error")
	}

	expectedValue := "127.0.0.1:9055"

	os.Setenv(envListeingAddr, expectedValue)

	value, err := getListeningAddr()
	if err != nil {
		t.Error("getListeningAddr should not return an error")
	}

	if value != expectedValue {
		t.Errorf("expected %s, actual %s", expectedValue, value)
	}

	teardown()
}

func Test_GetInterval(t *testing.T) {
	setup()

	_, err := getInterval()
	if err == nil {
		t.Error("getInterval should return an error")
	}

	os.Setenv(envInterval, "15")

	value, err := getInterval()
	if err != nil {
		t.Error("getInterval should not return an error")
	}

	expectedValue := 15 * time.Second
	if value != expectedValue {
		t.Errorf("expected %s, actual %s", expectedValue, value)
	}

	teardown()
}
