package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	LoadEnv()

	host := MustGetEnv("DB_HOST")
	user := MustGetEnv("DB_USER")
	password := MustGetEnv("DB_PASSWORD")
	dbname := MustGetEnv("DB_NAME")
	port := MustGetEnv("DB_PORT")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Error trying to connect to the database:", err)
	}
	fmt.Println("Connection successful")

	return db
}
