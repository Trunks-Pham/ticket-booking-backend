// @title Ticket Booking API
// @version 1.0
// @description API cho ứng dụng đặt vé với các chức năng quản lý người dùng và chuyến bay.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

package main

import (
	"log"

	_ "github.com/Trunks-Pham/ticket-booking-backend/cmd/api/docs" // Local import of docs package
	"github.com/Trunks-Pham/ticket-booking-backend/internal/initialize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger" // Import Fiber Swagger middleware
)

func main() {
	// Tạo instance của Fiber
	app := fiber.New()

	// Thêm middleware ghi log request
	app.Use(logger.New())

	// Route cho tài liệu Swagger
	app.Get("/api/v1/swagger/*", swagger.HandlerDefault) // Endpoint for Swagger Docs

	// Khởi tạo các route, middleware và service khác
	initialize.Run(app)

	// Khởi động ứng dụng tại cổng 8080
	log.Fatal(app.Listen(":8080"))
}
