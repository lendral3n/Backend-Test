package handler

import (
	"lendra/features/user"
	"lendra/utils/middlewares"
	"lendra/utils/responses"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) RegisterUser(c *fiber.Ctx) error {
	newUser := UserRequest{}
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	if err := handler.userService.Create(userCore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error insert data. "+err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success insert user", nil))
}

func (handler *UserHandler) Login(c *fiber.Ctx) error {
	var reqData = LoginRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error bind data, data not valid", nil))
	}

	result, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error login. "+err.Error(), nil))
	}

	access_token, err := middlewares.CreateTokenLogin(int(result.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse("error creating access token. "+err.Error(), nil))
	}

	refresh_token, err := middlewares.CreateRefreshToken(int(result.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse("error creating refresh token. "+err.Error(), nil))
	}

	// Set the tokens as cookies on the client side
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		Expires:  time.Now().Add(24 * 7 * time.Hour), // 7 days
		HTTPOnly: true,
	})

	responseData := LoginResponseData{
		Nama:         result.Name,
		Role:         result.Role,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) GetUser(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	result, err := handler.userService.GetById(userIdLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse("error read data. "+err.Error(), nil))
	}

	userResult := CoreToResponse(*result)
	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success read data", userResult))
}

func (handler *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	var userData = UserRequest{}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := UpdateRequestToCoreUpdate(userData)
	if err := handler.userService.Update(userIdLogin, userCore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error update data. "+err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	if err := handler.userService.Delete(userIdLogin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse("error delete data. "+err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success delete data", nil))
}

func (handler *UserHandler) ChangePassword(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	var passwords = ChangePasswordRequest{}
	if err := c.BodyParser(&passwords); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error bind data. data not valid", nil))
	}

	if err := handler.userService.ChangePassword(userIdLogin, passwords.OldPassword, passwords.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error change password. "+err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success change password", nil))
}

func (handler *UserHandler) AdminCreateUser(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	newUser := UserRequest{}
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	if err := handler.userService.AdminCreateUser(userIdLogin, userCore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.WebResponse("error insert data. "+err.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success insert user", nil))
}

func (handler *UserHandler) AdminGetAllUsers(c *fiber.Ctx) error {
	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse(err.Error(), nil))
	}

	users, err := handler.userService.AdminGetAllUser(userIdLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.WebResponse(err.Error(), nil))
	}

	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, CoreToResponse(user))
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success get data", usersResponse))
}

func (handler *UserHandler) RefreshToken(c *fiber.Ctx) error {
	userID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid access token")
	}

	accessToken, err := middlewares.CreateTokenLogin(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to create access token")
	}

	refreshToken, err := middlewares.CreateRefreshToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to create refresh token")
	}

	// Set the tokens as cookies on the client side
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour), 
		HTTPOnly: true,
	})

	responseData := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success refresh token", responseData))
}

func (handler *UserHandler) Logout(c *fiber.Ctx) error {
	// Remove the token cookie
	c.ClearCookie("access_token")
	c.ClearCookie("refresh_token")

	return c.Status(fiber.StatusOK).JSON(responses.WebResponse("success logout", nil))
}
