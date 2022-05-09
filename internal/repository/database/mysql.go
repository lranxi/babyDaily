package database

import (
	"baby-daily-api/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

var DB = new(gorm.DB)

// NewMysql 连接mysql
func init() {
	conf := configs.Get().Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(formatLoggerLevel()),
	})
	if err != nil {
		panic(err)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(conf.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifeTime)
	DB = db
}

func Get() *gorm.DB {
	return DB
}

func formatLoggerLevel() logger.LogLevel {

	switch strings.ToLower(configs.Get().Mysql.LoggerLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Silent
	}

}
