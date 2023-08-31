package main

import (
    fiber "github.com/gofiber/fiber/v2"
    postgres "github.com/Tulenien/Avito-backend-trainee-assignment-2023/db"
)

func someFunc(val int) int {
	return val + 2
}

func main() {
 postgres.ConnectDb()

 app := fiber.New()

 app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, world")
 })
 
 app.Listen(":3000")
}