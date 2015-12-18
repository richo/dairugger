package dairugger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/user"
)

type Client struct {
	client http.Client
}

type VoltronResponse struct {
	Type   string          `json: "type"`
	Status string          `json: "status"`
	Data   json.RawMessage `json: "data"`
}

func NewClient() Client {
	usr, _ := user.Current()

	return Client{
		// There's no obvious way to make this DTRT so we just throw everything
		// they want on the floor and connect to voltron. Fuck you golang.
		http.Client{
			// omfg welcome to randomly adding *'s and &'s until it compiles.
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					sock := fmt.Sprintf("%s/.voltron/sock", usr.HomeDir)
					return net.Dial("unix", sock)
				},
			},
		},
	}
}

func (c *Client) Get(route string) (*VoltronResponse, error) {
	// Our broken dialer just throws the hostname on the ground, but we'll call
	// it voltron for prettiness reasons
	path := fmt.Sprintf("http://voltron/api/%s", route)
	resp, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var data VoltronResponse
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Status != "success" {
		return nil, errors.New("Api request failed")
	}

	return &data, nil
}
