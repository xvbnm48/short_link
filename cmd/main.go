package main

import (
	"api-service/internal/db"
	"api-service/internal/repository"
	"api-service/internal/service"
	"api-service/internal/usecase"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT_SERVER")
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
	v1 := app.Group("/v1")
	v1.Post("/shorten", repoUsecase.CreateShortLink)
	v1.Get("/shorten/links", repoUsecase.GetAllLink)
	v1.Get("/shorten/:id", repoUsecase.GetShortLink)
	v1.Get("/:shortCode", repoUsecase.GetOriginalURL)

	log.Info("Starting server on :", port)
	app.Get("/hello", HelloHandler)
	app.Get("/health", HealthHandler)
	if err != app.Listen(":"+port) {
		log.Error("Failed to start server:", err)
		return
	}
	log.Info("Server is running on port:", port)
	log.Info("Server started successfully")
}
