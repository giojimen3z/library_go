package config

import (
	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := ConnectDB()
	db.AutoMigrate(&model.Author{})

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorApplication := application.NewAuthorApp(authorService)
	authorController := controller.NewAuthorController(authorApplication)

	r := gin.Default()
	r.POST("/authors", authorController.Create)
	r.GET("/authors", authorController.GetAll)
	r.GET("/author/:id", authorController.GetById)
	r.PUT("/authors/:id", authorController.Update)
	r.Run(":8080")
}
