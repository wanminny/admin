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
		rbac := middleware.RbacHandler()
		jwt := middleware.JWTAuthMiddleware()
		logger := middleware.DefaultLogger()
		group := global.Eng.Group("/v1/user").Use(rbac)
		user := handlers.NewUser()
		group.POST("/", user.Reg)
		group.GET("/:id", user.Login).Use(logger)
		group.PUT("/:id", user.Modify).Use(middleware.JWTAuthMiddleware())
		group.DELETE("/:id", user.Delete).Use(middleware.JWTAuthMiddleware())
		group.GET("/all", user.All).Use(jwt)
	} else {
		klog.Infoln("eng is nil")
	}
}
