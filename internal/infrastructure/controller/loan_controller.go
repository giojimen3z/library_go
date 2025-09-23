package controller

import (
	"library/internal/application"
	"library/internal/domain/model"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoanController struct {
	app application.LoanUseCaseInteface
}

func NewLoanController(app application.LoanUseCaseInteface) *LoanController {
	return &LoanController{app}
}

func (c *LoanController) Create(ctx *gin.Context) {
	var loan model.Loan

	if err := ctx.ShouldBindJSON(&loan); err != nil {
		slog.Error("Error trying to convert loan", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.app.CreateLoanUseCase(ctx.Request.Context(), &loan); err != nil {
		slog.Error("Error trying to create loan", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (c *LoanController) GetAll(ctx *gin.Context) {
	loans, err := c.app.GetLoansUseCase(ctx.Request.Context())
	if err != nil {
		slog.Error("Error trying to find loans", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}

	loan, err := c.app.GetLoanUseCase(ctx.Request.Context(), id)
	if err != nil {
		slog.Error("Error trying to get loan", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetByMemberId(ctx *gin.Context) {
	idParam := ctx.Param("memberId")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Member ID"})
		return
	}

	loan, err := c.app.GetLoanByMemeberIdUseCase(ctx.Request.Context(), id)
	if err != nil {
		slog.Error("Error trying to get loan", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}
func (c *LoanController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		slog.Error("Error trying to get parameter", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var patch model.Loan
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		slog.Error("Error trying to convert loan", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.app.UpdateLoanUseCase(ctx.Request.Context(), id, &patch)
	if err != nil {
		slog.Error("Error trying to update loan", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
