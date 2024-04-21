package middlewares

import (
	"fmt"
	"lendra/app/config"
	"strings"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware JWT untuk Fiber
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Try to get token from Authorization header
		authHeader := c.Get("Authorization")
		tokenString := ""
		if authHeader != "" {
			tokenString = strings.Split(authHeader, " ")[1]
		} else {
			// If Authorization header is empty, try to get token from cookie
			tokenString = c.Cookies("access_token")
			if tokenString == "" {
				return c.Status(fiber.StatusUnauthorized).SendString("Missing Authorization header")
			}
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWT_SECRET), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("user", claims)
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		return c.Next()
	}
}

// Generate token JWT
func CreateTokenLogin(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

// Ekstrak ID pengguna dari token JWT
func ExtractTokenUserId(c *fiber.Ctx) (int, error) {
	user := c.Locals("user")
	if user == nil {
		return 0, nil
	}
	claims := user.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return int(userId), nil
}

// Generate refresh token
func CreateRefreshToken(userId int) (string, error) {
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["userId"] = userId
	// Set longer expiration for refresh token
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	return refreshToken.SignedString([]byte(config.JWT_SECRET))
}

// Logout User
func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
