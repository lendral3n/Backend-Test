package service

import (
	"lendra/features/user"
	"lendra/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	hash := new(mocks.HashInterface)
	srv := New(repo, hash)

	inputData := user.User{
		ID:           1,
		Name:         "budi",
		Email:        "budi@gmail.com",
		Password:     "12345678",
		Role:         "user",
	}

	t.Run("Success Create User", func(t *testing.T) {
		hash.On("HashPassword", mock.AnythingOfType("string")).Return("hashedPass", nil).Once()
		repo.On("Insert", mock.Anything).Return(nil).Once()
		inputData.Role = ""
		err := srv.Create(inputData)
		inputData.Role = "renter"

		assert.NoError(t, err)
		hash.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()
		invalidInput := user.User{}
		err := srv.Create(invalidInput)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required")
		repo.AssertNotCalled(t, "Insert")
	})

	t.Run("Hash Password Error", func(t *testing.T) {
		hash.On("HashPassword", mock.Anything).Return("", errors.New("hash error")).Once()
		err := srv.Create(inputData)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error hash password")
		repo.AssertNotCalled(t, "Insert")
	})

}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	hash := new(mocks.HashInterface)

	returnData := user.User{
		ID:           1,
		Name:         "budi",
		Email:        "budi@gmail.com",
		Password:     "12345678",
		Role:         "user",
	}

	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&returnData, nil).Once()
		srv := New(repo, hash)
		result, err := srv.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData.Name, result.Name)
		assert.Equal(t, returnData.Email, result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		repo.On("SelectById", 1).Return(nil, errors.New("user not found")).Once()
		srv := New(repo, hash)
		result, err := srv.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	hash := new(mocks.HashInterface)
	userService := New(repo, hash)

	input := user.UserUpdate{
		Name:         "budi",
		Email:        "budi@gmail.com",
	}

	t.Run("Validation Error", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(nil).Once()
		invalidInput := user.User{}
		err := userService.Create(invalidInput)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required")
		repo.AssertNotCalled(t, "Update")
	})

	t.Run("invalid user id", func(t *testing.T) {
		err := userService.Update(0, input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid id")
	})

	t.Run("success update", func(t *testing.T) {
		userId := 3
		repo.On("Update", userId, input, mock.Anything).Return(nil).Once()
		err := userService.Update(userId, input)

		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	hash := new(mocks.HashInterface)
	userService := New(repo, hash)

	t.Run("invalid user id", func(t *testing.T) {
		err := userService.Delete(0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid id")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", 1).Return(nil).Once()

		err := userService.Delete(1)

		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserDataInterface)
	hash := new(mocks.HashInterface)
	userService := New(repo, hash)

	inputLogin := user.User{
		ID:       1,
		Email:    "updated@gmail.com",
		Password: "newpassword",
	}

	t.Run("empty email and password", func(t *testing.T) {
		_, err := userService.Login("", "")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email dan password wajib diisi")
	})

	t.Run("empty email", func(t *testing.T) {
		_, err := userService.Login("", "password")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email wajib diisi")
	})

	t.Run("empty password", func(t *testing.T) {
		_, err := userService.Login("email@gmail.com", "")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password wajib diisi")
	})

	t.Run("password not match", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(&inputLogin, nil).Once()
		hash.On("CheckPasswordHash", inputLogin.Password, inputLogin.Password).Return(false).Once()

		_, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password tidak sesuai")
	})

	t.Run("error on userData.Login", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(nil, errors.New("some error")).Once()

		_, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "some error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(&inputLogin, nil).Once()
		hash.On("CheckPasswordHash", inputLogin.Password, inputLogin.Password).Return(true).Once()

		data, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.NoError(t, err)
		assert.Equal(t, inputLogin, *data)
	})
}

func TestChangePassword(t *testing.T) {
    repo := new(mocks.UserDataInterface)
    hash := new(mocks.HashInterface)
    userService := New(repo, hash)

     t.Run("invalid user id", func(t *testing.T) {
        repo.On("SelectById", 0).Return(nil, errors.New("invalid id")).Once()
        hash.On("HashPassword", "newPassword").Return("hashedNewPassword", nil).Once()
        repo.On("ChangePassword", 0, "oldPassword", "hashedNewPassword").Return(errors.New("invalid id")).Once()
        err := userService.ChangePassword(0, "oldPassword", "newPassword")

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "invalid id")
    })

    t.Run("empty old password", func(t *testing.T) {
        repo.On("SelectById", 1).Return(&user.User{Password: "hashedOldPassword"}, nil).Once()
        err := userService.ChangePassword(1, "", "newPassword")

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "please input current password")
    })

    t.Run("empty new password", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&user.User{Password: "hashedOldPassword"}, nil).Once()
        err := userService.ChangePassword(1, "oldPassword", "")

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "please input new password")
    })

    t.Run("old password not match", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&user.User{Password: "hashedOldPassword"}, nil).Once()
        hash.On("CheckPasswordHash", "hashedOldPassword", "oldPassword").Return(false).Once()
        hash.On("HashPassword", "newPassword").Return("hashedNewPassword", nil).Once()
        repo.On("ChangePassword", 1, "oldPassword", "hashedNewPassword").Return(errors.New("current password not match")).Once() 
        err := userService.ChangePassword(1, "oldPassword", "newPassword")

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "current password not match")
    })

    t.Run("new password same as old password", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&user.User{Password: "hashedOldPassword"}, nil).Once()
        hash.On("CheckPasswordHash", "hashedOldPassword", "oldPassword").Return(true).Once()
        hash.On("HashPassword", "newPassword").Return("hashedNewPassword", nil).Once() 
        hash.On("HashPassword", "newPassword").Return("hashedNewPassword", nil).Once() 
        repo.On("ChangePassword", 1, "oldPassword", "hashedNewPassword").Return(errors.New("password cannot be the same")).Once()
        err := userService.ChangePassword(1, "oldPassword", "newPassword")

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "password cannot be the same")
    })

    t.Run("success", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&user.User{Password: "hashedOldPassword"}, nil).Once()
        hash.On("CheckPasswordHash", "hashedOldPassword", "oldPassword").Return(true).Once()
        hash.On("CheckPasswordHash", "hashedOldPassword", "newPassword").Return(false).Once()
        hash.On("HashPassword", "newPassword").Return("hashedNewPassword", nil).Once()
        repo.On("ChangePassword", 1, "oldPassword", "hashedNewPassword").Return(nil).Once()

        err := userService.ChangePassword(1, "oldPassword", "newPassword")

        assert.NoError(t, err)
    })
}

func TestAdminCreateUser(t *testing.T) {
    repo := new(mocks.UserDataInterface)
    hash := new(mocks.HashInterface)
    userService := New(repo, hash)

    t.Run("admin create user", func(t *testing.T) {
        repo.On("SelectById", 1).Return(&user.User{Role: "admin"}, nil).Once()
        hash.On("HashPassword", "password").Return("hashedPassword", nil).Once()
        repo.On("Insert", mock.Anything).Return(nil).Once()

        err := userService.AdminCreateUser(1, user.User{Password: "password"})

        assert.NoError(t, err)
    })

    t.Run("non-admin cannot create user", func(t *testing.T) {
        repo.On("SelectById", 2).Return(&user.User{Role: "user"}, nil).Once()

        err := userService.AdminCreateUser(2, user.User{Password: "password"})

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "anda bukan admin")
    })
}

func TestAdminGetAllUser(t *testing.T) {
    repo := new(mocks.UserDataInterface)
    hash := new(mocks.HashInterface)
    userService := New(repo, hash)

    t.Run("admin get all users", func(t *testing.T) {
        repo.On("SelectById", 1).Return(&user.User{Role: "admin"}, nil).Once()
        repo.On("AdminGetAllUser").Return([]user.User{}, nil).Once()

        _, err := userService.AdminGetAllUser(1)

        assert.NoError(t, err)
    })

    t.Run("non-admin cannot get all users", func(t *testing.T) {
        repo.On("SelectById", 2).Return(&user.User{Role: "user"}, nil).Once()

        _, err := userService.AdminGetAllUser(2)

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "anda bukan admin")
    })
}
