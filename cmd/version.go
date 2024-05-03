package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd 代表 "version" 命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of admin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("your-command version 1.0.0")
	},
}
