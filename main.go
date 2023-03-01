package main

import (
	"mini-project-product/db"
	"mini-project-product/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.DBInit()

	app := fiber.New()

	route.GetUsers(app)

	app.Listen(":8000")
}
