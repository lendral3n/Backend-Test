package service

import (
	"errors"
	"lendra/features/user"
	"lendra/utils/encrypts"
	"sync"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
	m           sync.Map
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.User) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("error hash password")
		}
		input.Password = hashedPass
	}

	err := service.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(userId int) (*user.User, error) {
	result, err := service.userData.SelectById(userId)
	return result, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userId int, input user.UserUpdate) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	if userId <= 0 {
		return errors.New("invalid id")
	}

	err := service.userData.Update(userId, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userId int) error {
	if userId <= 0 {
		return errors.New("invalid id")
	}
	err := service.userData.Delete(userId)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.User, err error) {
	if email == "" && password == "" {
		return nil,errors.New("email dan password wajib diisi")
	}
	if email == "" {
		return nil,errors.New("email wajib diisi")
	}
	if password == "" {
		return nil, errors.New("password wajib diisi")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil,err
	}

	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil,errors.New("password tidak sesuai")
	}

	return data, nil
}

// ChangePassword implements user.UserServiceInterface.
func (service *userService) ChangePassword(userId int, oldPassword, newPassword string) error {
	if oldPassword == "" {
		return errors.New("please input current password")
	}

	if newPassword == "" {
		return errors.New("please input new password")
	}

	hashedNewPass, _ := service.hashService.HashPassword(newPassword)

	err := service.userData.ChangePassword(userId, oldPassword, hashedNewPass)
	return err
}

// AdminCreateUser implements user.UserServiceInterface.
func (service *userService) AdminCreateUser(userId int, input user.User) error {
	//  cek role admin
	user, _ := service.userData.SelectById(userId)
	if user.Role != "admin" {
		return errors.New("anda bukan admin")
	}

	if input.Password != "" {
		hashedPass, _ := service.hashService.HashPassword(input.Password)
		input.Password = hashedPass
	}

	err := service.userData.Insert(input)
	return err
}

// AdminGetAllUser implements user.UserServiceInterface.
func (service *userService) AdminGetAllUser(userId int) ([]user.User, error) {
	user, err := service.userData.SelectById(userId)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.New("anda bukan admin")
	}

	users, err := service.userData.AdminGetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}
