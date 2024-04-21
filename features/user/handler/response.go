package handler

import "lendra/features/user"

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
}

type LoginResponseData struct {
    Nama         string `json:"nama"`
    Role         string `json:"role"`
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

func CoreToResponse(data user.User) UserResponse {
	var result = UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}
	return result
}
