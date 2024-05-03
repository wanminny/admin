package global

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var Eng *gin.Engine
var once sync.Once

//var mu sync.Mutex

func Init() {
	once.Do(func() {

		if Eng == nil {
			//mu.Lock()
			// 初始化 Viper 配置
			//initConfig()
			//mode := viper.GetBool("http.port")
			//klog.Infoln(mode)
			//if mode {
			//	gin.SetMode(gin.DebugMode)
			//} else {
			//	gin.SetMode(gin.ReleaseMode)
			//}
			gin.SetMode(gin.DebugMode)
			Eng = gin.Default()
			//klog.Infoln(Eng)
			//mu.Unlock()
		}
	})
}
