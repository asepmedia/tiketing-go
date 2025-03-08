package service

import (
	event "codereview/api/ticketing/modules/event/repository"
	"codereview/api/ticketing/modules/ticket/repository"
	"codereview/internal/domain"
	"errors"
	"sync"
)

type TicketService struct {
	eventRepo  *event.EventRepo
	ticketRepo *repository.TicketRepo
	mu         sync.Mutex
}

func NewTicketService(eventRepo *event.EventRepo, ticketRepo *repository.TicketRepo) *TicketService {
	return &TicketService{
		eventRepo:  eventRepo,
		ticketRepo: ticketRepo,
	}
}

func (s *TicketService) PurchaseTicket(eventID, userID int64) (*domain.Ticket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Ambil data event
	event, err := s.eventRepo.GetEvent(eventID)
	if err != nil {
		return nil, err
	}

	// Pastikan masih ada tiket tersedia
	if event.Sold >= event.Capacity {
		return nil, errors.New("tickets sold out")
	}

	// Buat tiket baru
	ticket := &domain.Ticket{
		EventID: eventID,
		UserID:  userID,
	}

	createdTicket, err := s.ticketRepo.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}

	// Update jumlah tiket terjual
	event.Sold++

	err = s.eventRepo.UpdateEvent(event)
	if err != nil {
		return nil, err
	}

	return createdTicket, nil
}
