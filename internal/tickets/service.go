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

// Todo
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
	tickets, _ := s.GetTotalTickets(ctx, destination)
	totalTickets, _ := s.r.GetAll(ctx)
	avg := len(tickets) / len(totalTickets)

	return avg, nil
}
