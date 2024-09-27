package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authService services.IAuthService,
	flightRepository repositories.IFlightRepository,
	ticketRepository repositories.ITicketRepository,
) {
	api := app.Group("/api")

	authRoute(api.Group("/auth"), authService)

	flightRoutes(api.Group("/flight"), flightRepository)

	ticketRoutes(api.Group("/tickets"), ticketRepository)
}
