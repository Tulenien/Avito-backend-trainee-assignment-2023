package main

import (
    fiber "github.com/gofiber/fiber/v2"
    "github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/database"
    //"github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/repositories"
)

func someFunc(val int) int {
	return val + 2
}

func main() {
 database.ConnectDB()
 //userRepository := repositories.NewUsersRepository(database.DB.DB)

 app := fiber.New()

 setupRoutes(app)
 
 app.Listen(":3000")
}