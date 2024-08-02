package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitializeDatabaseConnection() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s",
		viper.Get("database.user"),
		viper.Get("database.pass"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.databaseName"),
		"user_management",
	)
	newDB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("cannot connect to database!")
	}
	postgreDB, err := newDB.DB()
	postgreDB.SetMaxOpenConns(200)
	postgreDB.SetMaxIdleConns(100)
	fmt.Printf("connecting to %s\n", connectionString)
	if err != nil {
		panic("can't connect to database!")
	}

	DB = newDB
	fmt.Printf("Connected to database: %s\n", connectionString)
}
