package ibc

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"net/url"
	"io/ioutil"
)

type IBCMonitor struct {
	host string
	urlTarget string

	supplyTemperature string
	returnTemperature string
	targetTemperature string
	inletPressure string
	deltaPressure string
	status string

	mbh int32
}

func NewMonitor(host string) IBCMonitor {
	return IBCMonitor{
		host: host,
		urlTarget: fmt.Sprintf("http://%s/%s", host, TARGET_URL),
	}
}

func (i *IBCMonitor) GetExtendedDetails() (*IBCResponseExtDetail, error) {
	request_payload := IBCRequest{
		Object_no: 100,
		Object_request: IBC_BOILER_EXT_DETAIL_DATA,
		Boiler_no: 0,
	}

	payload_bytes, err := json.Marshal(request_payload)
	if err != nil {
		log.Print("Failed to serialize request payload", err)
		return nil, err
	}

	log.Print("Request Payload: ", request_payload)
	target := fmt.Sprintf("%s?json=%s", i.urlTarget, url.QueryEscape(string(payload_bytes[:])))
	log.Print("Requesting ", target)
	resp, err := http.Get(target)
	if err != nil {
		log.Print("HTTP Request to boiler failed", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Print("Unexpected response code: ", resp.StatusCode)
		log.Print("Unexpected response: ", resp.Status)
		return nil, fmt.Errorf("Unexpected HTTP Response: %s", resp.Status)
	}

	responseBody := IBCResponseExtDetail{}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Failed to read response")
		return nil, err
	}

	err = json.Unmarshal(content, &responseBody)
	if err != nil {
		log.Print("Failed to unmarshal response")
		return nil, err
	}

	// IBC Uses some weird multipliers, scale things down before returning the object
	responseBody.SupplyT = responseBody.SupplyT/4
	responseBody.ReturnT = responseBody.ReturnT/4
	responseBody.TargetT = responseBody.TargetT/4

	return &responseBody, nil
}
