package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes thiết lập các route chính cho ứng dụng
// Các route bao gồm auth, flight và ticket với các hành động tương ứng.
// @param app là Fiber App chính cho ứng dụng.
// @param authService là service xử lý các tác vụ xác thực người dùng.
// @param flightRepository là repository cho các chức năng liên quan đến chuyến bay.
// @param ticketRepository là repository cho các chức năng liên quan đến vé.
func SetupRoutes(
	app *fiber.App,
	authService models.IAuthService,
	flightRepository models.IFlightRepository,
	ticketRepository models.ITicketRepository,
) {
	// Tạo nhóm route chung với prefix /api
	api := app.Group("/api")

	// Định nghĩa các route cho module xác thực
	// Các route này sẽ nằm trong group /api/auth
	authRoute(api.Group("/auth"), authService)

	// Định nghĩa các route cho module chuyến bay
	// Các route này sẽ nằm trong group /api/flight
	flightRoutes(api.Group("/flight"), flightRepository)

	// Định nghĩa các route cho module vé
	// Các route này sẽ nằm trong group /api/tickets
	ticketRoutes(api.Group("/tickets"), ticketRepository)
}
