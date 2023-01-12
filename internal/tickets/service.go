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
	s.r.GetAll(ctx)
	// filtrar por pais
	lista := []
	return lista, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error) {
	lista, _ := GetTotalTickets(ctx, destination)
	avg := 0
	return avg, nil
}
