package main

import (
	"codereview/api/ticketing/modules/event"
	"codereview/api/ticketing/modules/ticket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// register event router

	api := app.Group("/api")
	v1 := api.Group("/v1")

	eventModule := event.NewEventRouter(v1)
	ticket.NewTicketRouter(v1, eventModule.Repo)

	// start listen app with port 3000
	app.Listen(":3000")
}
