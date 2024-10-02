package initialize

import (
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories/implement"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/routes"
	implement2 "github.com/Trunks-Pham/ticket-booking-backend/internal/services/implement"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run() {
	LoadConfig()
	InitPostgreSql()

	app := fiber.New(fiber.Config{
		AppName:      "FightBooking",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                   // Hoặc bạn có thể chỉ định danh sách các origin được phép
		AllowMethods: "GET,POST,PUT,DELETE", // Các phương thức được phép
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Server is Ok!",
		})
	})

	// Repository
	flightRepository := implement.NewFlightRepository()
	ticketRepository := implement.NewTicketRepository()
	authRepository := implement.NewAuthRepository()
	bookingRepository := implement.NewBookingHistoryRepository()

	// Service
	authService := implement2.NewAuthService(authRepository)
	bookingService := implement2.NewBookingService(bookingRepository, ticketRepository)

	// Setup routes
	routes.SetupRoutes(app, authService, flightRepository, ticketRepository, bookingService, bookingRepository)

	app.Listen(fmt.Sprintf(":%v", global.Config.Server.Port))
}
