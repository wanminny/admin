package routers

import (
	"github.com/wanminny/admin/pkg/global"
	"github.com/wanminny/admin/pkg/handlers"
	"github.com/wanminny/admin/pkg/middleware"
	"k8s.io/klog/v2"
)

func init() {
	global.Init()
	if global.Eng != nil {
		group := global.Eng.Group("/v1/user").Use(middleware.DefaultLogger())
		user := handlers.NewUser()
		group.POST("/:id", user.Reg)
		group.GET("/:id", user.GetUser)
	} else {
		klog.Infoln("eng is nil")
	}
}
