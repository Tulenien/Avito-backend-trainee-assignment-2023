package main

import (
    "os"

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
 
 err := app.Listen(":3000")
 if err != nil {
    os.Exit(1)
 }
 os.Exit(0)
}