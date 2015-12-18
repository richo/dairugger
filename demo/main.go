package main

import (
	"github.com/richo/dairugger"
)

func main() {
	client := dairugger.NewClient()
	client.Get()
}
