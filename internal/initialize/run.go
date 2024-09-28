package initialize

import (
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/routes"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	LoadConfig()
	InitPostgreSql()

	app := fiber.New(fiber.Config{
		AppName:      "FightBooking",
		ServerHeader: "Fiber",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Server is Ok!",
		})
	})

	// Repository
	flightRepository := repositories.NewFlightRepository()
	ticketRepository := repositories.NewTicketRepository()
	authRepository := repositories.NewAuthRepository()

	// Service
	authService := services.NewAuthService(authRepository)

	// Setup routes
	routes.SetupRoutes(app, authService, flightRepository, ticketRepository)

	app.Listen(fmt.Sprintf(":%v", global.Config.Server.Port))
}
