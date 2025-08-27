package app

import (
	"library/internal/infrastructure/controller"

	"github.com/gin-gonic/gin"
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

	registerAuthorRoutes(base, h.Author)
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
