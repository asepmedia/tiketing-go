package repository

import (
	"codereview/internal/domain"
	"errors"
	"sync"
)

type EventRepo struct {
	mu     sync.Mutex
	events *[]domain.Event
}

func NewEventRepo() *EventRepo {
	return &EventRepo{
		events: &[]domain.Event{},
	}
}

func (r *EventRepo) CreateEvent(event *domain.Event) (*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	event.ID = int64(len(*r.events) + 1)
	*r.events = append(*r.events, *event)
	return event, nil
}

func (r *EventRepo) GetAllEvent() (*[]domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.events, nil
}

func (r *EventRepo) GetEvent(id int64) (*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var event *domain.Event

	for _, e := range *r.events {
		if e.ID == id {
			event = &e
		}
	}

	if event == nil {
		return nil, errors.New("event not found")
	}

	return event, nil
}

func (r *EventRepo) UpdateEvent(event *domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Akses slice langsung melalui pointer
	events := *r.events

	for i := range events {
		if events[i].ID == event.ID {
			// Update langsung di dalam slice
			events[i].Sold = event.Sold
			events[i].Name = event.Name
			events[i].Capacity = event.Capacity
			return nil
		}
	}

	return errors.New("event not found")
}
