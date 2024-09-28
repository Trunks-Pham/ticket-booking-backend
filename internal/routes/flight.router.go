package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/middlewares"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// flightRoutes thiết lập các route cho việc quản lý chuyến bay (flight).
// @param route là Fiber Router để định nghĩa các đường dẫn.
// @param flightRepository là repository xử lý các tác vụ liên quan đến chuyến bay.
func flightRoutes(route fiber.Router, flightRepository models.IFlightRepository) {
	flightController := controllers.NewFlightController(flightRepository)

	privateRoutes := route.Use(middlewares.AuthProtected())

	// @Summary Lấy danh sách các chuyến bay
	// @Description Lấy danh sách tất cả các chuyến bay
	// @Tags Flight
	// @Produce json
	// @Success 200 {array} models.Flight "Danh sách chuyến bay"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/flight [get]
	privateRoutes.Get("/", flightController.GetMany)

	// @Summary Tạo chuyến bay mới
	// @Description Thêm một chuyến bay mới vào hệ thống
	// @Tags Flight
	// @Accept json
	// @Produce json
	// @Param flight body models.Flight true "Thông tin chuyến bay"
	// @Success 201 {object} models.Flight
	// @Failure 400 {object} fiber.Map "Bad Request"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/flight [post]
	privateRoutes.Post("/", flightController.CreateOne)

	// @Summary Lấy thông tin chuyến bay
	// @Description Lấy thông tin chi tiết của một chuyến bay bằng flightId
	// @Tags Flight
	// @Produce json
	// @Param flightId path string true "ID chuyến bay"
	// @Success 200 {object} models.Flight
	// @Failure 404 {object} fiber.Map "Not Found"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/flight/{flightId} [get]
	privateRoutes.Get("/:flightId", flightController.GetOne)

	// @Summary Cập nhật chuyến bay
	// @Description Cập nhật thông tin của một chuyến bay
	// @Tags Flight
	// @Accept json
	// @Produce json
	// @Param flightId path string true "ID chuyến bay"
	// @Param flight body models.Flight true "Thông tin chuyến bay cập nhật"
	// @Success 200 {object} models.Flight
	// @Failure 400 {object} fiber.Map "Bad Request"
	// @Failure 404 {object} fiber.Map "Not Found"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/flight/{flightId} [put]
	privateRoutes.Put("/:flightId", flightController.UpdateOne)

	// @Summary Xóa chuyến bay
	// @Description Xóa một chuyến bay khỏi hệ thống
	// @Tags Flight
	// @Produce json
	// @Param flightId path string true "ID chuyến bay"
	// @Success 200 {object} fiber.Map "Deleted"
	// @Failure 404 {object} fiber.Map "Not Found"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/flight/{flightId} [delete]
	privateRoutes.Delete("/:flightId", flightController.DeleteOne)
}
