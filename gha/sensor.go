package gha

import (
	"../ibc"
)

type GHA struct {
	secret string
	sensor string
	host string
}

func NewGHA(host string, secret string, sensor string) *GHA {
	return &GHA {
		host: host,
		secret: secret,
		sensor: sensor,
	}
}

func (gha *GHA) createSensorValue(subtype string, value string) *boilerPayload {
	return &boilerPayload{
		sensorPayload.Sensor: gha.sensor,
		sensorPayload.Secret: gha.secret,
		SensorType: "boiler_temperature",

		Subtype: subtype,
		Value: value,
	}
}

