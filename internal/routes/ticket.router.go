package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

func ticketRoutes(route fiber.Router, repository repositories.ITicketRepository) {
	ticketController := controllers.NewTicketController(repository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	privateRoutes.Get("/", ticketController.GetMany)
	privateRoutes.Get("/:ticketId", ticketController.GetOne)

	privateRoutes.Post("/", middlewares.CheckManagerRole(), ticketController.CreateOne)
	privateRoutes.Put("/:ticketId", middlewares.CheckManagerRole(), ticketController.UpdateOne)
	privateRoutes.Delete("/:ticketId", middlewares.CheckManagerRole(), ticketController.DeleteOne)
}
