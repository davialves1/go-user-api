package models

import (
	"github.com/google/uuid"
	"time"
)

type TechnicalUser struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Gid       string    `json:"gid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
