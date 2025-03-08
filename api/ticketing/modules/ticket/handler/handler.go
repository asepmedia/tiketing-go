package handler

import (
	"codereview/api/ticketing/modules/ticket/service"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct {
	ticketService *service.TicketService
}

func NewTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) PurchaseTicket(c *fiber.Ctx) error {
	var req struct {
		EventID int64 `json:"event_id"`
		UserID  int64 `json:"user_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	ticket, err := h.ticketService.PurchaseTicket(req.EventID, req.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(ticket)
}
