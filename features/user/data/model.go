package data

import (
	"lendra/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string
}

func UserToModel(input user.User) User {
	return User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}
}

func UserToModelUpdate(input user.UserUpdate) User {
	return User{
		Name:  input.Name,
		Email: input.Email,
	}
}

func (u User) ModelToUser() user.User {
	return user.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
