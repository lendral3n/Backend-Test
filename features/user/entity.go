package user

import (
	"time"
)

type User struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdate struct {
	Name  string 
	Email string 
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input User) error
	SelectById(userId int) (*User, error)
	Update(userId int, input UserUpdate) error
	Delete(userId int) error
	Login(email, password string) (data *User, err error)
	ChangePassword(userId int, oldPassword, newPassword string) error
	AdminGetAllUser() ([]User, error)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input User) error
	GetById(userId int) (*User, error)
	Update(userId int, input UserUpdate) error
	Delete(userId int) error
	Login(email, password string) (data *User, err error)
	ChangePassword(userId int, oldPassword, newPassword string) error
	AdminCreateUser(userId int, input User) error
	AdminGetAllUser(userId int) ([]User, error)
}
