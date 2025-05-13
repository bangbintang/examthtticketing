package service

import (
	"errors"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"

	"github.com/google/uuid"
)

type TicketService struct {
	ticketRepo *repository.TicketRepository
}

func NewTicketService(ticketRepo *repository.TicketRepository) *TicketService {
	return &TicketService{ticketRepo: ticketRepo}
}

func (s *TicketService) PurchaseTicket(ticket *models.Ticket) (*models.Ticket, error) {
	if ticket.UserID == uuid.Nil {
		return nil, errors.New("user ID tidak boleh kosong")
	}
	if ticket.EventID == 0 {
		return nil, errors.New("event ID tidak boleh kosong")
	}
	if err := s.ticketRepo.Create(ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *TicketService) GetTicketsByUser(userID string) ([]models.Ticket, error) {
	return s.ticketRepo.FindByUserID(userID)
}

func (s *TicketService) GetTicketsByEvent(eventID int) ([]models.Ticket, error) {
	return s.ticketRepo.FindByEventID(eventID)
}

func (s *TicketService) CancelTicket(ticketID int) error {
	return s.ticketRepo.Delete(ticketID)
}
