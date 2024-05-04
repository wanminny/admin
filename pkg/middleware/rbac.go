package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/wanminny/admin/pkg/global"
	"github.com/wanminny/admin/pkg/util"
	"k8s.io/klog/v2"
	"sync"
)

var (
	syncedCachedEnforcer *casbin.Enforcer
	once                 sync.Once
)

func Casbin() *casbin.Enforcer {
	//once.Do(func()
	{
		klog.Infoln("=============================Casbin======")
		klog.Infoln(global.Db)
		adapter, err := gormadapter.NewAdapterByDB(global.Db)
		if err != nil {
			klog.Infoln("适配数据库失败请检查casbin表是否为InnoDB引擎!")
			//return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow)) # Passes auth if any of the policies allows
		
		[role_definition]
		g = _, _    # user to roles
		
		[matchers]
		m = g(r.sub, p.sub) && r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch3(r.obj, p.obj)) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			klog.Infoln("字符串加载模型失败!")
			//return
		}
		syncedCachedEnforcer, err = casbin.NewEnforcer(m, adapter)
		//syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
		return syncedCachedEnforcer
	}

}

func RbacHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的PATH
		r := util.NewResponse()
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := "admin"
		klog.Infoln(sub, obj, act)
		e := Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			util.SetFailed(c, r, errors.New("权限不足或者数据库没有该功能Policy"))
			c.Abort()
			return
		}
		c.Next()
	}
}
