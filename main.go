package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalkrsharma/go-blogs/database"
)

func init() {
	database.ConnectDB()

}

func main() {

	postgresDB, err := database.DBConn.DB()

	if err != nil {
		panic("DB connection error")
	}

	defer postgresDB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})

	})

	app.Listen(":5000")

}
