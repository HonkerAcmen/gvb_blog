package core

import (
	"gvb_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置mysql,取消连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	if global.Config.System.Env == "debug" {
		// 打印所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Error("数据库连接错误!")
	}
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(10)               // 设置最大空闲连接数
	sqldb.SetMaxOpenConns(100)              // 设置最多可容纳
	sqldb.SetConnMaxLifetime(time.Hour * 4) // 设置最大复用时间，最多不超过mysql的wait_timeout
	global.Log.Info("数据库连接成功!")
	return db
}
