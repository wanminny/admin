package main

import "log"

func main() {

	// Cobra 也会将 flag 作为命令行参数
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")

	// 初始化 Viper 配置
	initConfig()

	// 添加子命令到 rootCmd
	rootCmd.AddCommand(versionCmd, startCmd)

	// 执行 rootCmd，它将调用上面的 Run 函数
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
