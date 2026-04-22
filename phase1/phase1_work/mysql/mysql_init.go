package mysql

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局DB对象
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// MySQL DSN 格式：user:pass@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:616069433@tcp(127.0.0.1:3306)/local_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志级别 INFO 会打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取 sqlDB 失败：%v", err)
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 空闲连接
	sqlDB.SetMaxOpenConns(100)          // 最大连接
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最长存活时间

	// 赋值给全局DB
	DB = db

	log.Println("✅ 数据库初始化成功")
}
