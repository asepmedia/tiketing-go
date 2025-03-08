package event

import (
	"codereview/api/ticketing/modules/event/handler"
	"codereview/api/ticketing/modules/event/repository"
	"codereview/api/ticketing/modules/event/service"
	"github.com/gofiber/fiber/v2"
)

type EventRouterResponse struct {
	Repo    *repository.EventRepo
	Service *service.EventService
}

func NewEventRouter(router fiber.Router) *EventRouterResponse {
	eventRepo := repository.NewEventRepo()
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	route := router.Group("/events")
	route.Get("/", eventHandler.GetAllEvent)
	route.Post("/", eventHandler.CreateEvent)

	return &EventRouterResponse{
		Repo:    eventRepo,
		Service: eventService,
	}
}
