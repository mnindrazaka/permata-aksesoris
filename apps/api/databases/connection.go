package databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	databaseName string
	username     string
	password     string
	host         string
	port         int
}

func NewDBConnection(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.username, config.password, config.host, config.port, config.databaseName)
	con, err := gorm.Open(mysql.Open(dsn))
	return con, err
}

func NewMainDBConnection() (*gorm.DB, error) {
	return NewDBConnection(Config{databaseName: "permata_aksesoris", host: "127.0.0.1", port: 3306, username: "root", password: "roottoor"})
}

func NewTestDBConnection() (*gorm.DB, error) {
	return NewDBConnection(Config{databaseName: "permata_aksesoris_test", host: "127.0.0.1", port: 3306, username: "root", password: "roottoor"})
}
