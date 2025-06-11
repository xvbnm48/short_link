package main

import (
	"api-service/internal/db"
	"api-service/internal/repository"
	"api-service/internal/service"
	"api-service/internal/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HealthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello world",
	})
}

func main() {
	fmt.Println("hello world!")
	app := fiber.New()

	// to database connection
	db, err := db.NewDB()
	if err != nil {
		log.Error("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	repoLink := repository.NewLinkRepository(db)
	repoService := service.NewLinkService(repoLink)
	repoUsecase := usecase.NewLinkUseCase(repoService)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	app.Post("/shorten", repoUsecase.CreateShortLink)
	app.Get("/shorten/:id", repoUsecase.GetShortLink)
	// app.Get("/shorten/:id", repoUsecase.)

	log.Info("Starting server on :3000")
	app.Get("/hello", HelloHandler)
	app.Get("/health", HealthHandler)
	app.Listen(":3000")
}
