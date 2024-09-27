package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

func flightRoutes(route fiber.Router, flightRepository repositories.IFlightRepository) {
	flightController := controllers.NewFlightController(flightRepository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	privateRoutes.Get("/", flightController.GetMany)
	privateRoutes.Get("/:flightId", flightController.GetOne)

	privateRoutes.Post("/", middlewares.CheckManagerRole(), flightController.CreateOne)
	privateRoutes.Put("/:flightId", middlewares.CheckManagerRole(), flightController.UpdateOne)
	privateRoutes.Delete("/:flightId", middlewares.CheckManagerRole(), flightController.DeleteOne)
}
