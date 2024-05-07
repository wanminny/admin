package global

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/wanminny/admin/pkg/gormx"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
	"sync"
)

var Db *gorm.DB
var ormOnce sync.Once

func LoadOrm() *gorm.DB {
	klog.Infoln(Db)
	if Db != nil {
		return Db // 避免重复初始化
	} else {
		ormOnce.Do(func() {
			config := gormx.Config{}
			db := viper.Get("db")
			b, err := json.Marshal(db)
			if err != nil {
				klog.Errorln(err)
			}
			err = json.Unmarshal(b, &config)
			Db, err = gormx.New(config)
			klog.Infoln(config, Db)
			klog.Infof("%+v", config)
			if err != nil {
				klog.Infoln(err.Error())
			}
		})
		return Db
	}
}
