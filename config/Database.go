package config //config/Database.go

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB Orm
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig Settings of DB
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "db",
		Port:     3306,
		User:     "root",
		Password: "MySql2020!",
		DBName:   "testdb",
	}
	return &dbConfig
}

// DbURL Config
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
