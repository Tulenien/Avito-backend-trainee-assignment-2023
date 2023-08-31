package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/segments", handlers.CreateNewSegment)
	app.Get("/segments/:id", handlers.DeleteSegment)
	app.Get("/users/:id/segments/[]:name/segments/[]:name", handlers.UpdateUserSegments)
	app.Get("/users/:id/segments", handlers.GetUserSegments)
}