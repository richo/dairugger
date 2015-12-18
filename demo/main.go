package main

import (
	"fmt"
	"github.com/richo/dairugger"
)

func main() {
	client := dairugger.NewClient()

	targets, _ := client.GetTargets()
	fmt.Println(targets)

	registers, _ := client.GetX64Registers()
	fmt.Println(registers)
}
