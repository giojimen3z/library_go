package config

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	LoadEnv()

	host := MustGetEnv("DB_HOST")
	user := MustGetEnv("DB_USER")
	password := MustGetEnv("DB_PASSWORD")
	dbname := MustGetEnv("DB_NAME")
	port := MustGetEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error("Error trying to connect to the database", "error", err)
		os.Exit(1)
	}

	slog.Info("Connection successful")

	return db
}
