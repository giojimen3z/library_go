package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load("cmd/api/app/infrastructure/config/dev.env")
	if err != nil {
		log.Fatalf("Error load .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // salida: consola
		logger.Config{
			SlowThreshold: time.Second, // consultas lentas > 1s
			LogLevel:      logger.Info, // nivel: Silent, Error, Warn, Info
			Colorful:      true,        // colores en consola
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
