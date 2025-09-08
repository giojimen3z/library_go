package controller

import (
	"log/slog"
	"net/http"

	"library/internal/application"
	"library/internal/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	app application.BookUseCaseInterface
}

func NewBookController(app application.BookUseCaseInterface) *BookController {
	return &BookController{app}
}

func (c *BookController) Create(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		slog.Error("Error trying to convert book", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.app.CreateBookUseCase(&book)
	if err != nil {
		slog.Error("Error trying to create book", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

func (c *BookController) GetAll(ctx *gin.Context) {
	books, err := c.app.GetBooksUseCase()
	if err != nil {
		slog.Error("Error trying to find books", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, err := c.app.GetBookByIdUseCase(id)
	if err != nil {
		slog.Error("Error trying to get book", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var patch model.Book
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		slog.Error("Error trying to convert book", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.app.UpdateBookUseCase(id, &patch)
	if err != nil {
		slog.Error("Error trying to update book", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
