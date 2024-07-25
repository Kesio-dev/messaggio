package stats

import "github.com/gofiber/fiber/v2"

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c Controller) GetAll(ctx *fiber.Ctx) error {
	stats, err := c.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"stats": stats,
	})
}
