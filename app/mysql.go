package app

import (
	"belajar-redis/helper"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB() *gorm.DB {
	sqlDB, err := sql.Open("mysql", "root:aryahmph@tcp(mysql:3306)/belajar_redis?charset=utf8&parseTime=True&loc=Local")
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))

	return db
}
