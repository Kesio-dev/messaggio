package message

import (
	"github.com/gofiber/fiber/v2"
	"messaggio/internal/kafka"
	"messaggio/internal/repositories"
)

func SetupRouter(app *fiber.App, messageRepo repositories.MessageRepositoryInterface, kafkaProducer *kafka.Producer) {
	service := NewService(messageRepo, kafkaProducer)
	controller := NewController(service)

	api := app.Group("/message")

	api.Post("/", controller.SaveMessage)
}
