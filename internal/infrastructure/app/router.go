package app

import (
	"github.com/gin-gonic/gin"

	"library/internal/infrastructure/controller"
)

type Handlers struct {
	Author *controller.AuthorController
}

func NewHandlers(author *controller.AuthorController) Handlers {
	return Handlers{
		Author: author,
	}
}

func RegisterRoutes(r *gin.Engine, h Handlers) {
	base := r.Group("/api/v1/library")

	healthRoutes(r)
	registerAuthorRoutes(base, h.Author)
}

func healthRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "pong",
		})
	})
}

func registerAuthorRoutes(group *gin.RouterGroup, c *controller.AuthorController) {
	if c == nil {
		return
	}
	authors := group.Group("/authors")
	authors.POST("", c.Create)
	authors.GET("", c.GetAll)
	authors.GET("/:id", c.GetById)
	authors.PUT("/:id", c.Update)
}
