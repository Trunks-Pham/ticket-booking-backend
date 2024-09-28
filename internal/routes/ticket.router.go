package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ticketRoutes thiết lập các route cho việc quản lý vé (tickets).
// @param route là Fiber Router để định nghĩa các đường dẫn.
// @param repository là repository xử lý các tác vụ liên quan đến vé.
func ticketRoutes(route fiber.Router, repository models.ITicketRepository) {
	ticketController := controllers.NewTicketController(repository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	// @Summary Lấy danh sách vé
	// @Description Lấy danh sách tất cả các vé
	// @Tags Tickets
	// @Produce json
	// @Success 200 {array} models.Ticket "Danh sách vé"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/tickets [get]
	privateRoutes.Get("/", ticketController.GetMany)

	// @Summary Tạo vé mới
	// @Description Thêm một vé mới vào hệ thống
	// @Tags Tickets
	// @Accept json
	// @Produce json
	// @Param ticket body models.Ticket true "Thông tin vé"
	// @Success 201 {object} models.Ticket
	// @Failure 400 {object} fiber.Map "Bad Request"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/tickets [post]
	privateRoutes.Post("/", ticketController.CreateOne)

	// @Summary Lấy thông tin vé
	// @Description Lấy thông tin chi tiết của một vé bằng ticketId
	// @Tags Tickets
	// @Produce json
	// @Param ticketId path string true "ID vé"
	// @Success 200 {object} models.Ticket
	// @Failure 404 {object} fiber.Map "Not Found"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/tickets/{ticketId} [get]
	privateRoutes.Get("/:ticketId", ticketController.GetOne)

	// @Summary Cập nhật thông tin vé
	// @Description Cập nhật thông tin của một vé
	// @Tags Tickets
	// @Accept json
	// @Produce json
	// @Param ticketId path string true "ID vé"
	// @Param ticket body models.Ticket true "Thông tin vé cập nhật"
	// @Success 200 {object} models.Ticket
	// @Failure 400 {object} fiber.Map "Bad Request"
	// @Failure 404 {object} fiber.Map "Not Found"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/tickets/{ticketId} [put]
	privateRoutes.Put("/:ticketId", ticketController.UpdateOne)
}
