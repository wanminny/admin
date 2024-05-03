package global

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/wanminny/admin/pkg/gormx"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
)

var Db *gorm.DB

func LoadOrm() *gorm.DB {
	config := gormx.Config{}
	db := viper.Get("db")
	b, err := json.Marshal(db)
	if err != nil {
		klog.Errorln(err)
	}
	err = json.Unmarshal(b, &config)
	Db, err = gormx.New(config)
	klog.Infof("%+v", config)
	if err != nil {
		klog.Infoln(err.Error())
	}
	return Db
}
