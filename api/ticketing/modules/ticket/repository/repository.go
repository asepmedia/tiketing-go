package repository

import (
	"codereview/internal/domain"
	"errors"
	"sync"
)

type TicketRepo struct {
	mu      sync.Mutex
	tickets map[int64][]*domain.Ticket
}

func NewTicketRepo() *TicketRepo {
	return &TicketRepo{
		tickets: make(map[int64][]*domain.Ticket),
	}
}

func (r *TicketRepo) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticket.ID = int64(len(r.tickets[ticket.EventID]) + 1)
	r.tickets[ticket.EventID] = append(r.tickets[ticket.EventID], ticket)
	return ticket, nil
}

func (r *TicketRepo) GetTicketsByEvent(eventID int64) ([]*domain.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tickets, exists := r.tickets[eventID]
	if !exists {
		return nil, errors.New("no tickets found for this event")
	}
	return tickets, nil
}
