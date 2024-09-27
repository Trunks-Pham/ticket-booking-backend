package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authService models.IAuthService,
	flightRepository models.IFlightRepository,
	ticketRepository models.ITicketRepository,
) {

	api := app.Group("/api")

	authRoute(api.Group("/auth"), authService)

	flightRoutes(api.Group("/flight"), flightRepository)

	ticketRoutes(api.Group("/tickets"), ticketRepository)
}
