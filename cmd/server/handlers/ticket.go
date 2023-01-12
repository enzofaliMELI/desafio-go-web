package handler

import (
	"net/http"

	tickets "github.com/bootcamp-go/desafio-go-web/internal/tickets"
	gin "github.com/gin-gonic/gin"
)

type Ticket struct {
	t tickets.Service
}

func NewTicket(t tickets.Service) *Ticket {
	return &Ticket{t: t}
}

func (s *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.t.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.t.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
