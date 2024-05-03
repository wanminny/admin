package routers

import (
	"github.com/wanminny/admin/pkg/global"
	"github.com/wanminny/admin/pkg/handlers"
	"github.com/wanminny/admin/pkg/middleware"
	"k8s.io/klog/v2"
)

func Init() {
	global.Init()
	global.LoadOrm()
	if global.Eng != nil {
		group := global.Eng.Group("/v1/user").Use(middleware.DefaultLogger())
		user := handlers.NewUser()
		group.POST("/", user.Reg)
		group.GET("/:id", user.Login)
		group.PUT("/:id", user.Modify).Use(middleware.JWTAuthMiddleware())
		group.DELETE("/:id", user.Delete).Use(middleware.JWTAuthMiddleware())
		group.GET("/", user.All).Use(middleware.JWTAuthMiddleware())
	} else {
		klog.Infoln("eng is nil")
	}
}
