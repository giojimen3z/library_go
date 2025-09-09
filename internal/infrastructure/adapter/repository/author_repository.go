package repository

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"library/internal/domain/model"
	"library/internal/domain/port"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) port.AuthorPort {
	return &AuthorRepositoryImpl{db: db}
}

func (r *AuthorRepositoryImpl) Save(ctx context.Context, author *model.Author) error {
	err := r.db.WithContext(ctx).Create(author).Error
	if err != nil {
		return err
	}
	slog.Info("Author save successful", "author", author)
	return nil
}

func (r *AuthorRepositoryImpl) FindAll(ctx context.Context) ([]model.Author, error) {
	var authors []model.Author
	err := r.db.WithContext(ctx).Find(&authors).Error
	if err != nil {
		slog.Error("Failed to find authors", "error", err)
		return nil, err
	}

	return authors, err
}

func (r *AuthorRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*model.Author, error) {
	var author model.Author
	err := r.db.WithContext(ctx).First(&author, id).Error
	if err != nil {
		slog.Error("Failed to find author", "error", err)
		return nil, err
	}
	return &author, err
}

func (r *AuthorRepositoryImpl) Update(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error) {
	if err := r.db.WithContext(ctx).Model(&model.Author{}).Where("id = ?", id).Updates(patch).Error; err != nil {
		slog.Error("Failed to update author", "error", err)
		return nil, err
	}
	return nil, nil
}
