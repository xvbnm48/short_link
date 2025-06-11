package usecase

import (
	"api-service/internal/model"
	"api-service/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LinkUseCase interface {
	CreateShortLink(ctx *fiber.Ctx) error
	GetShortLink(ctx *fiber.Ctx) error
}

type linkUseCase struct {
	service service.LinkService
}

// GetShortLink implements LinkUseCase.
func (l *linkUseCase) GetShortLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	link, err := l.service.GetShortLink(idInt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Link not found",
			"error":   err.Error(),
		})
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
			"error":   err.Error(),
		})
	}

	id, err := l.service.CreateShortLink(LinkReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create short link",
			"error":   err.Error(),
		})
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
