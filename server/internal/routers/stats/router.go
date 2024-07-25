package stats

import (
	"github.com/gofiber/fiber/v2"
	"messaggio/internal/repositories"
)

func SetupRouter(app *fiber.App, messageRepo repositories.MessageRepositoryInterface) {
	service := NewService(messageRepo)
	controller := NewController(service)

	api := app.Group("/stats")

	api.Get("/", controller.GetAll)
}
