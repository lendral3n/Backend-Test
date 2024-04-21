package data

import (
	"errors"
	"lendra/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.User) error {
	dataGorm := UserToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(userId int) (*user.User, error) {
	var userDataGorm User
	tx := repo.db.First(&userDataGorm, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := userDataGorm.ModelToUser()
	return &result, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(userId int, input user.UserUpdate) error {
	dataGorm := UserToModelUpdate(input)
	tx := repo.db.Model(&User{}).Where("id = ?", userId).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(userId int) error {
	tx := repo.db.Delete(&User{}, userId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.User, err error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		// return nil, tx.Error
		return nil, errors.New(" Invalid email or password")
	}
	result := userGorm.ModelToUser()
	return &result, nil
}

// ChangePassword implements user.UserDataInterface.
func (repo *userQuery) ChangePassword(userId int, oldPassword, newPassword string) error {
	var userGorm User
	userGorm.Password = newPassword
	tx := repo.db.Model(&User{}).Where("id = ?", userId).Updates(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// AdminGetAllUser implements user.UserDataInterface.
func (repo *userQuery) AdminGetAllUser() ([]user.User, error) {
	var usersGorm []User

	tx := repo.db.Find(&usersGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var users []user.User
	for _, userGorm := range usersGorm {
		users = append(users, userGorm.ModelToUser())
	}
	return users, nil
}
