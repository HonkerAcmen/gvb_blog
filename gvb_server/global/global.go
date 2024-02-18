package global

import (
	"gvb_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// 用于保存配置文件
	Config *config.Config

	// 数据库接口
	DB *gorm.DB

	// log接口
	Log *logrus.Logger
)
