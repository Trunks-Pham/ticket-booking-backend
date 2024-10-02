// @title Ticket Booking API
// @version 1.0
// @description API cho ứng dụng đặt vé với các chức năng quản lý người dùng và chuyến bay.
// @termsOfService http://yourapi.com/terms/
// @contact.name Hỗ Trợ API
// @contact.email support@yourapi.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

package main

import (
	"log"

	_ "github.com/Trunks-Pham/ticket-booking-backend/cmd/api/docs" // Local import of docs package
	"github.com/Trunks-Pham/ticket-booking-backend/internal/initialize"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger" // Import Fiber Swagger middleware
)

// @title Ticket Booking API Documentation
// @version 1.0
// @description Đây là API đặt vé máy bay
// @host localhost:8080
// @BasePath /api/v1

func main() {
	app := fiber.New()

	// Route for Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault) // Endpoint for Swagger Docs

	// Initialize and register routes, middlewares, etc.
	initialize.Run()

	// Start the application on port 8080
	log.Fatal(app.Listen(":8080"))
}
