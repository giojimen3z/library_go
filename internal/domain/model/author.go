package model

import (
	"time"
)

type Author struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `json:"first_name" gorm:"size:200;not null"`
	LastName  string    `json:"last_name" gorm:"size:200;not null"`
	Bio       *string   `json:"bio" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
