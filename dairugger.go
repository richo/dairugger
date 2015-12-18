package dairugger

import (
	"fmt"

	"net"
	"net/http"

	"os/user"
)

type Client struct {
	client http.Client
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

func (c *Client) Get() {
	// Our broken dialer just throws the hostname on the ground, but we'll call
	// it voltron for prettiness reasons
	resp, _ := c.client.Get("http://voltron/api/request")
	fmt.Println(resp)
}
