package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	startPath   = "/v1/s1/e5/actions/start"
	performPath = "/v1/s1/e5/actions/perform-turn"
)

type AttackPosition struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type PerformTurnRequest struct {
	Action         string         `json:"action"`
	AttackPosition AttackPosition `json:"attack_position"`
}

type PerformTurnResponse struct {
	PerformedAction string `json:"performed_action"`
	TurnsRemaining  int    `json:"turns_remaining"`
	TimeRemaining   int    `json:"time_remaining"`
	ActionResult    string `json:"action_result"`
	Message         string `json:"message"`
}

type ValidationError struct {
	Detail []struct {
		Loc  []interface{} `json:"loc"`
		Msg  string        `json:"msg"`
		Type string        `json:"type"`
	} `json:"detail"`
}

func (c *Client) startMission() error {
	resp, err := c.doRequest("POST", startPath, nil)
	if err != nil {
		return fmt.Errorf("failed to start mission: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response from start endpoint: %d", resp.StatusCode)
	}

	fmt.Println("Mission started successfully!")
	return nil
}

func (c *Client) performTurn(action string, position AttackPosition) error {
	request := PerformTurnRequest{
		Action:         action,
		AttackPosition: position,
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := c.doRequest("POST", performPath, bytes.NewReader(requestBody))
	if err != nil {
		return fmt.Errorf("failed to perform turn: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode == http.StatusOK {
		var successResponse PerformTurnResponse
		if err := json.Unmarshal(respBody, &successResponse); err != nil {
			return fmt.Errorf("failed to parse success response: %v", err)
		}
		fmt.Printf("Turn performed successfully! Response: %+v\n", successResponse)
		return nil
	} else if resp.StatusCode == http.StatusUnprocessableEntity {
		var validationError ValidationError
		if err := json.Unmarshal(respBody, &validationError); err != nil {
			return fmt.Errorf("failed to parse validation error response: %v", err)
		}
		fmt.Printf("Validation error: %+v\n", validationError)
		return fmt.Errorf("validation error occurred")
	} else {
		return fmt.Errorf("unexpected response from perform-turn endpoint: %d, body: %s", resp.StatusCode, respBody)
	}
}

func (c *Client) Valiant() error {
	if err := c.startMission(); err != nil {
		fmt.Printf("Error starting mission: %v\n", err)
		return err
	}

	enemyPosition := AttackPosition{X: "e", Y: 5}
	allyPosition := AttackPosition{X: "f", Y: 8}
	obstacles := []AttackPosition{
		{X: "e", Y: 2},
		{X: "h", Y: 3},
		{X: "e", Y: 5},
	}

	fmt.Printf("Enemy position: %v\n", enemyPosition)
	fmt.Printf("Ally position: %v\n", allyPosition)
	fmt.Printf("Obstacles: %v\n", obstacles)

	predictedPosition := AttackPosition{X: "f", Y: 6}
	fmt.Printf("Predicted enemy position: %v\n", predictedPosition)

	if err := c.performTurn("attack", predictedPosition); err != nil {
		fmt.Printf("Error performing turn: %v\n", err)
		return err
	}

	fmt.Println("Mission completed successfully!")

	return nil
}
