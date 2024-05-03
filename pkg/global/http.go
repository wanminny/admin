package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var Eng *gin.Engine
var once sync.Once
var mu sync.Mutex

func init() {
	fmt.Println(Eng)
	once.Do(func() {
		if Eng == nil {
			fmt.Println("init")
			//mu.Lock()
			Eng = gin.Default()
			fmt.Println(Eng)
			//mu.Unlock()
		}
	})
}
