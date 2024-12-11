package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Measurement struct {
	Distance string `json:"distance"`
	Time     string `json:"time"`
}

func (c *Client) GetMeasurement() (*Measurement, error) {
	var measurement Measurement
	var attempt int

	for {
		attempt++
		resp, err := c.doRequest("GET", "/v1/s1/e1/resources/measurement", nil)
		if err != nil {
			return nil, fmt.Errorf("error en la solicitud: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("error: %s", string(body))
		}

		if err := json.NewDecoder(resp.Body).Decode(&measurement); err != nil {
			return nil, fmt.Errorf("error al decodificar la respuesta: %w", err)
		}

		// Verificar si la medición es válida
		if measurement.Distance != "failed to measure, try again" && measurement.Time != "failed to measure, try again" {
			fmt.Printf("Medición obtenida en el intento %d: %+v\n", attempt, measurement)
			return &measurement, nil
		}

		fmt.Printf("Intento %d fallido, repitiendo...\n", attempt)
	}
}

func ParseDistance(distance string) (float64, error) {
	cleanDistance := strings.TrimSpace(strings.ReplaceAll(distance, "AU", ""))
	return strconv.ParseFloat(cleanDistance, 64)
}

func ParseTime(time string) (float64, error) {
	cleanTime := strings.TrimSpace(strings.ReplaceAll(time, "hours", ""))
	return strconv.ParseFloat(cleanTime, 64)
}

func CalculateOrbitalSpeed(measurement *Measurement) (int, error) {
	distanceAU, err := ParseDistance(measurement.Distance)
	if err != nil {
		return 0, fmt.Errorf("error al convertir Distance a float64: %w", err)
	}
	distanceKm := distanceAU * 149597870.7

	timeHours, err := ParseTime(measurement.Time)
	if err != nil {
		return 0, fmt.Errorf("error al convertir Time a float64: %w", err)
	}

	speed := distanceKm / timeHours
	return int(speed), nil
}
