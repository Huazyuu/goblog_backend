package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func init() {
	username := "root"   // 账号
	password := "yanjun" // 密码
	host := "127.0.0.1"  // 数据库地址，可以是Ip或者域名
	port := 3306         // 数据库端口
	Dbname := "gorm"     // 数据库名
	timeout := "10s"     // 连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{ // 命名策略
			TablePrefix:   "",    // 前缀
			SingularTable: false, // 表名单数
			NoLowerCase:   false, // 关闭大小写替换
		},
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 使用彩色打印
			}),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
}
func main() {

}
