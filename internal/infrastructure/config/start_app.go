package config

import (
	"library/cmd/api/app/domain/model"
	"library/cmd/api/app/domain/service"
	"library/cmd/api/app/infrastructure/adapter/repository"
	"library/cmd/api/app/infrastructure/controller"

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
