package builder

import (
	"time"

	"library/internal/domain/model"

	"github.com/google/uuid"
)

type AuthorBuilder struct {
	id        uuid.UUID
	firstName string
	lastName  string
	bio       string
	createdAt *time.Time
	updatedAt *time.Time
}

func NewAuthorBuilder() *AuthorBuilder {
	date := time.Now()
	return &AuthorBuilder{
		id:        uuid.New(),
		firstName: "John",
		lastName:  "Doe",
		bio:       "test",
		createdAt: &date,
		updatedAt: &date,
	}
}

func UpdateAuthorBuilder() *AuthorBuilder {
	date := time.Now()
	return &AuthorBuilder{
		id:        uuid.New(),
		firstName: "John",
		lastName:  "Doe",
		bio:       "test update",
		createdAt: &date,
		updatedAt: &date,
	}
}

func (b *AuthorBuilder) Build() *model.Author {
	return &model.Author{
		ID:        b.id,
		FirstName: b.firstName,
		LastName:  b.lastName,
		Bio:       b.bio,
		CreatedAt: *b.createdAt,
		UpdatedAt: *b.updatedAt,
	}
}
