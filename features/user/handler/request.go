package handler

import (
	"lendra/features/user"
)

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func RequestToCore(input UserRequest) user.User {
	return user.User{
		Name:             input.Name,
		Email:            input.Email,
		Password:         input.Password,
	}
}

func UpdateRequestToCore(input UserRequest) user.UserUpdate {
	return user.UserUpdate{
		Name:         input.Name,
		Email:        input.Email,
	}
}

func UpdateRequestToCoreUpdate(input UserRequest) user.UserUpdate {
	return user.UserUpdate{
		Name:         input.Name,
		Email:        input.Email,
	}
}
