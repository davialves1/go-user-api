package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Gid       string    `json:"gid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (User) New(email string, name string) User {
	return User{
		ID:        uuid.New(),
		Gid:       uuid.NewString(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UserDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u User) ToDto() UserDto {
	return UserDto{
		ID:    u.ID.String(),
		Email: u.Email,
		Name:  u.Name,
	}
}

type UserRequest struct {
	Name  string
	Email string
}
