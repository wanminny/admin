package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VERSION = "v0.0.1"

// versionCmd 代表 "version" 命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of admin",
	Run: func(cmd *cobra.Command, args []string) {
		//version := main.VERSION
		fmt.Println("admin version:", VERSION)
	},
}
