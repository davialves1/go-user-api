package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Gid            string    `json:"gid"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Name           string    `json:"name"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`
}

func (User) New(email, name, hashedPassword string) User {
	return User{
		ID:             uuid.New(),
		Gid:            uuid.NewString(),
		Email:          email,
		Name:           name,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (u User) ToDto() UserDto {
	return UserDto{
		ID:    u.ID.String(),
		Email: u.Email,
		Name:  u.Name,
		Gid:   u.Gid,
	}
}

type UserDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Gid   string `json:"gid"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type UserJWT struct {
	Email string
	Id    string
}
