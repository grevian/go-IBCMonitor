package gha

type sensorPayload struct {
	Sensor string
	Secret string
}

type boilerPayload struct {
	sensorPayload
	SensorType string `json:"type"`
	Subtype string
	Value string
}

