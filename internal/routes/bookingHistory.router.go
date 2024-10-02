package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func bookingRoutes(
	route fiber.Router,
	BookingService services.IBookingHistoryService,
	BookingRepository repositories.IBookingHistoryRepository,
) {
	bookingController := controllers.NewBookingHistoryController(BookingService, BookingRepository)

	privateRoute := route.Use(middlewares.AuthProtected())

	privateRoute.Get("/", bookingController.GetMany)
	privateRoute.Post("/", bookingController.BookTicket)
}
