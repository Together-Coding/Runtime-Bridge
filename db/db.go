package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var DB *gorm.DB
var dbErr error

func Initialize() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URI"), os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		panic("failed to connect database")
	}

	sqlDB, sqlDBErr := DB.DB()
	if sqlDBErr != nil {
		panic("failed to make database connection pool")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
