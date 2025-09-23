package app

import (
	"library/internal/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Author *controller.AuthorController
	Loan   *controller.LoanController
}

func NewHandlers(author *controller.AuthorController, loan *controller.LoanController) Handlers {
	return Handlers{
		Author: author,
		Loan:   loan,
	}
}

func RegisterRoutes(r *gin.Engine, h Handlers) {
	base := r.Group("/api/v1/library")

	healthRoutes(r)
	registerAuthorRoutes(base, h.Author)
	registerLoanRoutes(base, h.Loan)
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

func registerLoanRoutes(group *gin.RouterGroup, c *controller.LoanController) {
	if c == nil {
		return
	}
	loans := group.Group("/loans")
	loans.POST("", c.Create)
	loans.GET("", c.GetAll)
	loans.GET("/:id", c.GetById)
	loans.GET("/memberList/:memberId", c.GetByMemberId)
	loans.PUT("/:id", c.Update)
}
