package config

import (
	"hexagonal_base/cmd/api/app/domain/model"
	"hexagonal_base/cmd/api/app/domain/service"
	"hexagonal_base/cmd/api/app/infrastructure/adapter/repository"
	"hexagonal_base/cmd/api/app/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := ConnectDB()
	db.AutoMigrate(&model.Author{})

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	r := gin.Default()
	r.POST("/authors", authorController.Create)
	r.GET("/authors", authorController.GetAll)

	r.Run(":8080")
}
