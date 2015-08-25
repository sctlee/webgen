package main

import (
	"fmt"
	"os"

	"command"
)

func main() {
	fmt.Println("Hello Secret!")
	args := os.Args

	// if args == nil || len(args) < 2 {
	// 	fmt.Println("Usage")
	// 	return
	// }

	switch args[1] {
	case "init":
		command.Init()
	case "build":
		command.Build()
	case "push":
		command.Push()
	case "reset":
		command.Reset()
	default:
		fmt.Println("Usage")
	}
}
