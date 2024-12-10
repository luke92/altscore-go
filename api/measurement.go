package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Measurement struct {
	Distance float64 `json:"distance"`
	Time     float64 `json:"time"`
}

func (c *Client) GetMeasurement() (*Measurement, error) {
	resp, err := c.doRequest("GET", "/v1/s1/e1/resources/measurement", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: %s", string(body))
	}

	var measurement Measurement
	if err := json.NewDecoder(resp.Body).Decode(&measurement); err != nil {
		return nil, err
	}

	return &measurement, nil
}

func CalculateOrbitalSpeed(measurement *Measurement) int {
	distanceInKm := measurement.Distance * 149597870.7
	speed := distanceInKm / measurement.Time
	return int(speed)
}
