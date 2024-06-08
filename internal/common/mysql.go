package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wangzheng/brain/config"
	"wangzheng/brain/internal/entity"
)

func NewMysql() *gorm.DB {
	dsn := config.MySQLConfig().ConnectionString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&entity.Diary{})
	if err != nil {
		panic("failed to migrate database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	return db
}
