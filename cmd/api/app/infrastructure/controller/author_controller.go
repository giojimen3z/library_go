package controller

import (
	"hexagonal_base/cmd/api/app/domain/model"
	"hexagonal_base/cmd/api/app/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	service *service.AuthorService
}

func NewAuthorController(service *service.AuthorService) *AuthorController {
	return &AuthorController{service}
}

func (c *AuthorController) Create(ctx *gin.Context) {
	var author model.Author

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateAuthor(&author); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, author)
}

func (c *AuthorController) GetAll(ctx *gin.Context) {
	authors, err := c.service.GetAuthors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}
