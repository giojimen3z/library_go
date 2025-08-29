package controller

import (
	"log/slog"
	"net/http"

	"library/internal/application"
	"library/internal/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	if err := c.app.CreateAuthorUseCase(ctx.Request.Context(), &author); err != nil {
		slog.Error("Error trying to create author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, author)
}

func (c *AuthorController) GetAll(ctx *gin.Context) {
	authors, err := c.app.GetAuthorsUseCase(ctx.Request.Context())
	if err != nil {
		slog.Error("Error trying to find authors", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

func (c *AuthorController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := c.app.GetAuthorUseCase(ctx.Request.Context(), id)
	if err != nil {
		slog.Error("Error trying to get author", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (c *AuthorController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
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

	updated, err := c.app.UpdateAuthorUseCase(ctx.Request.Context(), id, &patch)
	if err != nil {
		slog.Error("Error trying to update author", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}
