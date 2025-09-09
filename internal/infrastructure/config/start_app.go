package config

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/infrastructure/app"
	"library/internal/infrastructure/controller"
)

func StartApp() {
	db := ConnectDB()
	if err := db.AutoMigrate(&model.Author{}); err != nil {
		slog.Error("Failed to migrate database", "error", err)
		os.Exit(1)
	}

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorApplication := application.NewAuthorUseCase(authorService)
	authorController := controller.NewAuthorController(authorApplication)

	r := gin.Default()
	handlers := app.NewHandlers(authorController)
	app.RegisterRoutes(r, handlers)

	if err := r.Run(":" + GetEnv("PORT", "8080")); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
