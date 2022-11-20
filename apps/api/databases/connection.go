package databases

import (
	"fmt"
	"os"
	"strconv"

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
	databaseName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return NewDBConnection(Config{databaseName: databaseName, host: host, port: port, username: username, password: password})
}

func NewTestDBConnection() (*gorm.DB, error) {
	databaseName := os.Getenv("DB_TEST_NAME")
	host := os.Getenv("DB_TEST_HOST")
	username := os.Getenv("DB_TEST_USERNAME")
	password := os.Getenv("DB_TEST_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_TEST_PORT"))
	if err != nil {
		return nil, err
	}
	return NewDBConnection(Config{databaseName: databaseName, host: host, port: port, username: username, password: password})
}
