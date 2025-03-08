package test

import (
	eventrepo "codereview/api/ticketing/modules/event/repository"
	"codereview/api/ticketing/modules/ticket/repository"
	"codereview/api/ticketing/modules/ticket/service"
	"codereview/internal/domain"
	"testing"
)

func TestPurchaseTicket(t *testing.T) {
	eventRepo := eventrepo.NewEventRepo()
	ticketRepo := repository.NewTicketRepo()
	ticketService := service.NewTicketService(eventRepo, ticketRepo)

	// Membuat event baru
	event, err := eventRepo.CreateEvent(&domain.Event{Name: "Concert", Capacity: 2})
	if err != nil {
		t.Fatalf("failed to create event: %v", err)
	}

	// Membeli tiket pertama
	ticket1, err := ticketService.PurchaseTicket(event.ID, 1)
	if err != nil {
		t.Fatalf("failed to purchase first ticket: %v", err)
	}

	// Membeli tiket kedua (harus berhasil)
	ticket2, err := ticketService.PurchaseTicket(event.ID, 2)
	if err != nil {
		t.Fatalf("failed to purchase second ticket: %v", err)
	}

	// Coba beli tiket ketiga (harus gagal karena kapasitas penuh)
	_, err = ticketService.PurchaseTicket(event.ID, 3)
	if err == nil {
		t.Fatalf("expected error due to full capacity, but got none")
	}

	// Memastikan tiket yang dibuat benar
	if ticket1.EventID != event.ID || ticket2.EventID != event.ID {
		t.Errorf("ticket event ID mismatch")
	}
}
