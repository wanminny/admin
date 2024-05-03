package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wanminny/admin/pkg/global"
	"k8s.io/klog/v2"
)

// startCmd 代表 "start" 命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Long:  `This will start the application using the provided configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		// 从配置文件中读取设置
		config := viper.Get("yourConfigSection")
		fmt.Printf("Starting the application with config: %+v\n", config)
		gin.SetMode(gin.ReleaseMode)
		klog.Infoln(gin.Mode())
		global.Eng.Run(":8090")
	},
}
