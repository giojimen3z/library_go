package config

import (
	"log/slog"
	"os"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/infrastructure/app"
	"library/internal/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	db := ConnectDB()
	if err := db.AutoMigrate(
		&model.Author{},
		&model.BookAuthor{},
		&model.BookCopy{},
		&model.Book{},
		&model.Loan{},
		&model.Member{},
		&model.Reservation{},
	); err != nil {
		slog.Error("Failed to migrate database", "error", err)
		os.Exit(1)
	}

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorApplication := application.NewAuthorUseCase(authorService)
	authorController := controller.NewAuthorController(authorApplication)

	loanRepo := repository.NewLoanRepository(db)
	memeberService := service.NewMemberService()
	bookCopyService := service.NewBookCopyService()
	loanService := service.NewLoanService(loanRepo, memeberService, bookCopyService)
	loanApplication := application.NewLoanUseCase(loanService)
	loanController := controller.NewLoanController(loanApplication)

	r := gin.Default()
	handlers := app.NewHandlers(authorController, loanController)
	app.RegisterRoutes(r, handlers)

	if err := r.Run(":" + GetEnv("PORT", "8080")); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
