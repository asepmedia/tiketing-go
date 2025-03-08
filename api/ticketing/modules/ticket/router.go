package ticket

import (
	event "codereview/api/ticketing/modules/event/repository"
	"codereview/api/ticketing/modules/ticket/handler"
	"codereview/api/ticketing/modules/ticket/repository"
	"codereview/api/ticketing/modules/ticket/service"
	"github.com/gofiber/fiber/v2"
)

func NewTicketRouter(router fiber.Router, eventRepo *event.EventRepo) {
	ticketRepo := repository.NewTicketRepo()
	ticketService := service.NewTicketService(eventRepo, ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)

	route := router.Group("/tickets")
	route.Post("/", ticketHandler.PurchaseTicket)
}
