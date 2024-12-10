package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Star struct {
	Id        string   `json:"id"`
	Resonance float64  `json:"resonance"`
	Position  Position `json:"position"`
}

func (c *Client) GetStars() ([]Star, error) {
	var allStars []Star
	page := 1
	baseURL := "/v1/s1/e2/resources/stars"

	for {
		fmt.Printf("Obteniendo estrellas de la pagina %d\n", page)

		url := fmt.Sprintf("%s?page=%d&sort-by=resonance&sort-direction=asc", baseURL, page)

		resp, err := c.doRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("error: %s", string(body))
		}

		var stars []Star
		if err := json.NewDecoder(resp.Body).Decode(&stars); err != nil {
			return nil, err
		}

		allStars = append(allStars, stars...)

		if len(stars) == 0 {
			break
		}

		page++
	}

	return allStars, nil
}

func CalculateAverageResonance(stars []Star) float64 {
	var sum float64
	for _, star := range stars {
		sum += star.Resonance
	}
	return sum / float64(len(stars))
}
