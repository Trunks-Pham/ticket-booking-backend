package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ticketRoutes(route fiber.Router, repository models.ITicketRepository) {
	ticketController := controllers.NewTicketController(repository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	privateRoutes.Get("/", ticketController.GetMany)
	privateRoutes.Post("/", ticketController.CreateOne)
	privateRoutes.Get("/:ticketId", ticketController.GetOne)
	privateRoutes.Put("/:ticketId", ticketController.UpdateOne)
	privateRoutes.Delete("/:ticketId", ticketController.DeleteOne)
}
