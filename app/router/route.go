package router

import (
	ud "lendra/features/user/data"
	uh "lendra/features/user/handler"
	us "lendra/features/user/service"

	pd "lendra/features/product/data"
	ph "lendra/features/product/handler"
	ps "lendra/features/product/service"
	"lendra/utils/encrypts"
	"lendra/utils/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, app *fiber.App) {
	hash := encrypts.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	productData := pd.New(db)
	productService := ps.New(productData)
	productHandlerAPI := ph.New(productService)

	// Middleware JWT 
	jwtMiddleware := middlewares.JWTMiddleware()

	// define routes/ endpoint AUTH
	app.Post("/login", userHandlerAPI.Login)
	app.Post("/logout", jwtMiddleware, userHandlerAPI.Logout)
	app.Post("/refresh-token", jwtMiddleware, userHandlerAPI.RefreshToken)

	// define routes/ endpoint ADMIN
	app.Post("/admin/user", jwtMiddleware, userHandlerAPI.AdminCreateUser)
	app.Get("/admin/user", jwtMiddleware, userHandlerAPI.AdminGetAllUsers)

	// define routes/ endpoint USER
	app.Post("/users", userHandlerAPI.RegisterUser)
	app.Get("/users", jwtMiddleware, userHandlerAPI.GetUser)
	app.Put("/users", jwtMiddleware, userHandlerAPI.UpdateUser)
	app.Delete("/users", jwtMiddleware, userHandlerAPI.DeleteUser)
	app.Put("/change-password", jwtMiddleware, userHandlerAPI.ChangePassword)

	// define routes/ endpoint PRODUCT
	app.Post("/product", jwtMiddleware, productHandlerAPI.CreateProduct)
	app.Get("/product", productHandlerAPI.GetAllProduct)
	app.Get("/product/:id", productHandlerAPI.GetProductById)
	app.Put("/product/:id", jwtMiddleware, productHandlerAPI.UpdateProduct)
	app.Delete("/product/:id", jwtMiddleware, productHandlerAPI.DeleteProduct)
}
