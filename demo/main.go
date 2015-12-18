package main

import (
	"fmt"
	"github.com/richo/dairugger"
)

func main() {
	client := dairugger.NewClient()
	targets, _ := client.GetTargets()

	fmt.Println(targets)
}
