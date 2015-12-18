package dairugger

import (
	"encoding/json"
)

type TargetsResponse struct {
	Targets []Target `json: "targets"`
}

type Target struct {
	Id    uint   `json: "id"`
	File  string `json: "file"`
	Arch  string `json: "arch"`
	State string `json: "state"`
}

func (c *Client) GetTargets() ([]Target, error) {
	resp, err := c.Get("targets")

	if err != nil {
		return nil, err
	}

	var data TargetsResponse
	err = json.Unmarshal(resp.Data, &data)

	if err != nil {
		return nil, err
	}

	return data.Targets, nil
}
