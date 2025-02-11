package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabase() (db *gorm.DB, err error) {
	dsn := viper.GetString("database.dsn")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
