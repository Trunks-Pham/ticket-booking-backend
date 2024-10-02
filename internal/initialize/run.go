package initialize

import (
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories/implement"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/routes"
	implement2 "github.com/Trunks-Pham/ticket-booking-backend/internal/services/implement"
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
	flightRepository := implement.NewFlightRepository()
	ticketRepository := implement.NewTicketRepository()
	authRepository := implement.NewAuthRepository()

	// Service
	authService := implement2.NewAuthService(authRepository)

	// Setup routes
	routes.SetupRoutes(app, authService, flightRepository, ticketRepository)

	app.Listen(fmt.Sprintf(":%v", global.Config.Server.Port))
}
