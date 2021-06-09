package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"api/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func NewDb() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	conn, err := gorm.Open(mysql.Open(connString()), &gorm.Config{
		Logger: newLogger(),
	})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	log.Println("success to connect db!")

	dbInstance = conn

	return dbInstance
}

func connString() string {
	dbConf := config.Conf.Db

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
}

func newLogger() logger.Interface {
	loggerConfig := logger.Config{
		SlowThreshold: time.Second,
		Colorful:      false,
	}
	if os.Getenv("GO_ENV") == "development" {
		loggerConfig.LogLevel = logger.Info
	} else {
		loggerConfig.LogLevel = logger.Silent
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		loggerConfig,
	)
}
