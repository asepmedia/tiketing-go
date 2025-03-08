package service

import (
	"codereview/api/ticketing/modules/event/repository"
	"codereview/internal/domain"
	"errors"
)

type EventService struct {
	eventRepo *repository.EventRepo
}

func NewEventService(eventRepo *repository.EventRepo) *EventService {
	return &EventService{eventRepo: eventRepo}
}

func (s *EventService) CreateEvent(name string, capacity int) (*domain.Event, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than zero")
	}

	event := &domain.Event{
		Name:     name,
		Capacity: capacity,
		Sold:     0,
	}

	return s.eventRepo.CreateEvent(event)
}

func (s *EventService) GetAllEvent() (*[]domain.Event, error) {
	return s.eventRepo.GetAllEvent()
}

func (s *EventService) GetEvent(id int64) (*domain.Event, error) {
	return s.eventRepo.GetEvent(id)
}
