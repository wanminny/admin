package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var configPath string

func initConfig() {
	if configPath != "" {
		// 从指定路径读取配置文件
		viper.SetConfigFile(configPath)
	} else {
		// 寻找默认的配置文件
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

// rootCmd 代表命令行程序的基础命令
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "A brief description of admin",
}

func Start() {
	// Cobra 也会将 flag 作为命令行参数
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")

	// 初始化 Viper 配置
	initConfig()

	// 添加子命令到 rootCmd
	rootCmd.AddCommand(startCmd, versionCmd)

	// 执行 rootCmd，它将调用上面的 Run 函数
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
