package main

import (
	"command"
	"fmt"
	"os"
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
	default:
		fmt.Println("Usage")
	}
}
