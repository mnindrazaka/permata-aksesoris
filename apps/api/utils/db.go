package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	dsn := "root:roottoor@tcp(127.0.0.1:3306)/permata_aksesoris?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := gorm.Open(mysql.Open(dsn))
	return con, err
}
