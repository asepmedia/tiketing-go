package service

import (
	eventrepo "codereview/api/ticketing/modules/event/repository"
	"codereview/internal/domain"
	"testing"
)

func TestCreateEVent(t *testing.T) {
	eventRepo := eventrepo.NewEventRepo()

	// Membuat event baru
	_, err := eventRepo.CreateEvent(&domain.Event{Name: "Concert", Capacity: 2})
	if err != nil {
		t.Fatalf("failed to create event: %v", err)
	}
}
