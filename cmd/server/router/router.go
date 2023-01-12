package router

import (
	//middleware "github.com/bootcamp-go/desafio-go-web/cmd/serve/middlewares"
	handler "github.com/bootcamp-go/desafio-go-web/cmd/server/handlers"
	domain "github.com/bootcamp-go/desafio-go-web/internal/domain"
	tickets "github.com/bootcamp-go/desafio-go-web/internal/tickets"
	gin "github.com/gin-gonic/gin"
)

type Router struct {
	db []domain.Ticket
	en *gin.Engine
}

func NewRouter(en *gin.Engine, db []domain.Ticket) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) MapRoutes() {
	// instances
	repository := tickets.NewRepository(r.db)
	service := tickets.NewService(repository)
	handler := handler.NewTicket(service)

	tickets := r.en.Group("/tickets")
	tickets.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
	tickets.GET("/getAverage/:dest", handler.AverageDestination())
}
