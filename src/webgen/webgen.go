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
	var template_repo string
	var has_parser bool
	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  `init`,
		Run: func(cmd *cobra.Command, args []string) {
			command.Init(template_repo, has_parser)
		},
	}

	cmdInit.Flags().StringVarP(&template_repo, "template", "t", command.DEFAULT_TMPL, "Choose one template")
	cmdInit.Flags().BoolVarP(&has_parser, "parser", "p", false, "if you want to have a simple tool to edit paper, set true")

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

	var cmdNew = &cobra.Command{
		Use:   "new",
		Short: "new",
		Long:  `new`,
		Run: func(cmd *cobra.Command, args []string) {
			command.New()
		},
	}

	var rootCmd = &cobra.Command{Use: "webgen"}
	rootCmd.AddCommand(cmdInit, cmdBuild, cmdPush, cmdReset, cmdNew)
	rootCmd.Execute()

}
