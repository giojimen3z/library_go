package config

import (
	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/infrastructure/controller"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := ConnectDB()
	if err := db.AutoMigrate(&model.Author{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorApplication := application.NewAuthorApp(authorService)
	authorController := controller.NewAuthorController(authorApplication)

	r := gin.Default()
	r.POST("/authors", authorController.Create)
	r.GET("/authors", authorController.GetAll)
	r.GET("/author/:id", authorController.GetById)
	r.PUT("/authors/:id", authorController.Update)
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
