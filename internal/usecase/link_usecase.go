package usecase

import (
	"api-service/internal/model"
	"api-service/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type LinkUseCase interface {
	CreateShortLink(ctx *fiber.Ctx) error
	GetShortLink(ctx *fiber.Ctx) error
	GetOriginalURL(ctx *fiber.Ctx) error
}

type linkUseCase struct {
	service service.LinkService
}

// GetOriginalURL implements LinkUseCase.
func (l *linkUseCase) GetOriginalURL(ctx *fiber.Ctx) error {
	shortCode := ctx.Params("shortCode")
	originalURL, err := l.service.GetOriginalURL(shortCode)
	if err != nil {
		log.Error("Error retrieving original URL:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Original URL not found",
		})
	}
	return ctx.Redirect(originalURL, fiber.StatusFound)
}

// GetShortLink implements LinkUseCase.
func (l *linkUseCase) GetShortLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	link, err := l.service.GetShortLink(idInt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Link not found",
		})
		log.Error("Error retrieving link:", err)
	}
	return ctx.JSON(fiber.Map{
		"message": "Link retrieved successfully",
		"data":    link,
		"status":  fiber.StatusOK,
		"code":    200,
	})
}

// CreateShortLink implements LinkUseCase.
func (l *linkUseCase) CreateShortLink(ctx *fiber.Ctx) error {
	LinkReq := model.LinkCreateRequest{}
	if err := ctx.BodyParser(&LinkReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		log.Error("Error parsing request body:", err)
	}

	id, err := l.service.CreateShortLink(LinkReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create short link",
		})
		log.Error("Error creating short link:", err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Short link created successfully",
		"data":    id,
		"status":  fiber.StatusCreated,
		"code":    201,
	})

}

func NewLinkUseCase(service service.LinkService) LinkUseCase {
	return &linkUseCase{service: service}
}
