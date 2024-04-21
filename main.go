package main

import (
	"lendra/app/config"
	"lendra/app/database"
	"lendra/app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.InitConfig()
	dbSql := database.InitDBMySQL(cfg)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost, https://l3n.my.id",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${method} ${path} ${latency}\n",
	}))

	router.InitRouter(dbSql, app)
	// Start server dan port
	log.Fatal(app.Listen(":8000"))
}
