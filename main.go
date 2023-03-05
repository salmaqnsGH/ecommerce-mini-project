package main

import (
	"log"
	"mini-project-product/db"
	"mini-project-product/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db.DBInit()

	app := fiber.New()

	route.GetUsers(app)

	app.Listen(":" + port)
}
