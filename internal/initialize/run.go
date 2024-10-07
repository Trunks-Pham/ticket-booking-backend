package initialize

import (
	"fmt"
	"log"

	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories/implement"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/routes"
	implement2 "github.com/Trunks-Pham/ticket-booking-backend/internal/services/implement"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Run initializes the application
func Run(app *fiber.App) {
	// Load configuration and connect to PostgreSQL
	LoadConfig()
	InitPostgreSql()

	// Add logger middleware for request logging
	app.Use(logger.New())

	// Basic health check route
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Server is running smoothly!",
		})
	})

	// Initialize repositories
	flightRepository := implement.NewFlightRepository()
	ticketRepository := implement.NewTicketRepository()
	authRepository := implement.NewAuthRepository()

	// Initialize services
	authService := implement2.NewAuthService(authRepository)

	// Setup routes and inject dependencies
	routes.SetupRoutes(app, authService, flightRepository, ticketRepository)

	// Start the application server
	port := fmt.Sprintf(":%v", global.Config.Server.Port)
	log.Printf(" Server running on http://localhost%s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf(" Failed to start server: %v", err)
	}
}
