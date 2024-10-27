package main

import (
	"github.com/Raviraj2000/chat-app/database"
	"github.com/Raviraj2000/chat-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Init()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
		AllowOrigins:     "http://localhost:8080",
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
