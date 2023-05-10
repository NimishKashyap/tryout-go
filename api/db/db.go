package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	db = Connect()

	return db
}

func Connect() *gorm.DB {
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")
	dbHost := viper.GetString("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", dbHost, username, dbName, password)

	sqlDB, err := sql.Open("pgx", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)

	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
