package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println()
		logrus.Warnln("未配置mysql 取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) // 打印所有
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的
	}
	global.MySqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   mysqlLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// global.Log.Fatalf(fmt.Sprintf("[%s] mysql 连接失败", dsn))
		logrus.Fatalf(fmt.Sprintf("[%s] mysql 连接失败", dsn))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.Mysql.MaxConnLifeTime))

	logrus.Infof(fmt.Sprintf("[%s:%d] mysql connects success", global.Config.Mysql.Host, global.Config.Mysql.Port))

	return db
}
