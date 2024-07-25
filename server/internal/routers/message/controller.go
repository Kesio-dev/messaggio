package message

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

type Message struct {
	Message string `json:"message"`
}

func (c Controller) SaveMessage(ctx *fiber.Ctx) error {
	p := new(Message)
	if err := ctx.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := c.service.SaveMessage(p.Message)
	if err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": p.Message,
		"status":  "success",
	})
}
