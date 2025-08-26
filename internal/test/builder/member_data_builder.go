package builder

import (
	"time"

	"library/internal/domain/model"

	"github.com/google/uuid"
)

type MemberBuilder struct {
	id        uuid.UUID
	fullName  string
	email     string
	phone     string
	isActive  bool
	createdAt *time.Time
	updatedAt *time.Time
}

func NewMemberBuilder() *MemberBuilder {
	date := time.Now()
	return &MemberBuilder{
		id:        uuid.New(),
		fullName:  "John Doe",
		email:     "john.doe@example.com",
		phone:     "+1-555-1234",
		isActive:  true,
		createdAt: &date,
		updatedAt: &date,
	}
}

func (b *MemberBuilder) Build() *model.Member {
	return &model.Member{
		ID:        b.id,
		FullName:  b.fullName,
		Email:     b.email,
		Phone:     b.phone,
		IsActive:  b.isActive,
		CreatedAt: *b.createdAt,
		UpdatedAt: *b.updatedAt,
	}
}
