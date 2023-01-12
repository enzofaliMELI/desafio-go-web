package tickets

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (int, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, _ := s.r.GetAll(ctx)

	var ticketFilter []domain.Ticket
	for _, ticket := range tickets {
		if ticket.Country == destination {
			ticketFilter = append(ticketFilter, ticket)
		}
	}

	return ticketFilter, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error) {
	tickets, _ := s.r.GetTicketByDestination(ctx, destination)
	totalTickets, _ := s.r.GetAll(ctx)
	avg := float64(len(tickets)) / float64(len(totalTickets)) * 100

	return int(avg), nil
}
