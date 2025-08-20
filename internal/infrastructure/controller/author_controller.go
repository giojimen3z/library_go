package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"library/internal/application"
	"library/internal/domain/model"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	app application.AuthorUseCaseInterface
}

func NewAuthorController(app application.AuthorUseCaseInterface) *AuthorController {
	return &AuthorController{app}
}

func (c *AuthorController) Create(ctx *gin.Context) {
	var author model.Author

	if err := ctx.ShouldBindJSON(&author); err != nil {
		slog.Error("Error trying to convert author", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.app.CreateAuthorUseCase(&author); err != nil {
		slog.Error("Error trying to create author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, author)
}

func (c *AuthorController) GetAll(ctx *gin.Context) {
	authors, err := c.app.GetAuthorsUseCase()
	if err != nil {
		slog.Error("Error trying to find authors", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

func (c *AuthorController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := c.app.GetAuthorUseCase(id)
	if err != nil {
		slog.Error("Error trying to get author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (c *AuthorController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var patch model.Author
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		slog.Error("Error trying to convert author", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.app.UpdateAuthorUseCase(id, &patch)
	if err != nil {
		slog.Error("Error trying to update author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}
