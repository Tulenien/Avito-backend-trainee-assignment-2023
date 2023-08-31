package main

import (
    fiber "github.com/gofiber/fiber/v2"
    "github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/db"
)

func someFunc(val int) int {
	return val + 2
}

func main() {
 db.ConnectDb()

 app := fiber.New()

 setupRoutes(app)
 
 app.Listen(":3000")
}