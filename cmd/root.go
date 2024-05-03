package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

//import _ "github.com/wanminny/admin/pkg/routers"

var configPath string

func initConfig() {
	if configPath != "" {
		// 从指定路径读取配置文件
		viper.SetConfigFile(configPath)
	} else {
		// 寻找默认的配置文件
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

// rootCmd 代表命令行程序的基础命令
var rootCmd = &cobra.Command{
	Use:   "your-command",
	Short: "A brief description of your command",
}
