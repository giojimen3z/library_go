package controller

import (
	"library/internal/application"
	"library/internal/domain/model"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	app *application.AuthorUseCase
}

func NewAuthorController(app *application.AuthorUseCase) *AuthorController {
	return &AuthorController{app}
}

func (c *AuthorController) Create(ctx *gin.Context) {
	var author model.Author

	if err := ctx.ShouldBindJSON(&author); err != nil {
		slog.Error("Error trying to convert author", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.app.CreateAuthor(&author); err != nil {
		slog.Error("Error trying to create author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, author)
}

func (c *AuthorController) GetAll(ctx *gin.Context) {
	authors, err := c.app.GetAuthors()
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

	author, err := c.app.GetAuthor(id)
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

	updated, err := c.app.UpdateAuthor(id, &patch)
	if err != nil {
		slog.Error("Error trying to update author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}
