package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func flightRoutes(route fiber.Router, flightRepository models.IFlightRepository) {
	flightController := controllers.NewFlightController(flightRepository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	privateRoutes.Get("/", flightController.GetMany)
	privateRoutes.Post("/", flightController.CreateOne)
	privateRoutes.Get("/:flightId", flightController.GetOne)
	privateRoutes.Put("/:flightId", flightController.UpdateOne)
	privateRoutes.Delete("/:flightId", flightController.DeleteOne)
}
