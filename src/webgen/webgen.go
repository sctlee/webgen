package main

import (
	"fmt"
	// "os"

	"command"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Hello Secret!")
	// args := os.Args

	// if args == nil || len(args) < 2 {
	// 	fmt.Println("Usage")
	// 	return
	// }

	// switch args[1] {
	// case "init":
	// 	command.Init()
	// case "build":
	// 	command.Build()
	// case "push":
	// 	command.Push()
	// case "reset":
	// 	command.Reset()
	// default:
	// 	fmt.Println("Usage")
	// }
	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  `init`,
		Run: func(cmd *cobra.Command, args []string) {
			command.Init()
		},
	}

	var cmdBuild = &cobra.Command{
		Use:   "build",
		Short: "build",
		Long:  `build`,
		Run: func(cmd *cobra.Command, args []string) {
			command.Build()
		},
	}

	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "push",
		Long:  `push`,
		Run: func(cmd *cobra.Command, args []string) {
			command.Push()
		},
	}

	var cmdReset = &cobra.Command{
		Use:   "reset",
		Short: "build",
		Long:  `build`,
		Run: func(cmd *cobra.Command, args []string) {
			command.Reset()
		},
	}
	// cmdHaha.Flags().IntVarP(&t, "type", "t", 1, "type hahaha")

	var rootCmd = &cobra.Command{Use: "webgen"}
	rootCmd.AddCommand(cmdInit, cmdBuild, cmdPush, cmdReset)
	rootCmd.Execute()

}
