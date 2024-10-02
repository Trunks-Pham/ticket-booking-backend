package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

// FlightController quản lý các yêu cầu liên quan đến chuyến bay
type FlightController struct {
	repository repositories.IFlightRepository
}

// Lấy danh sách các chuyến bay
// @Summary Lấy nhiều chuyến bay
// @Description Lấy danh sách các chuyến bay từ cơ sở dữ liệu
// @Tags Chuyến bay
// @Accept  json
// @Produce  json
// @Success 200 {object} fiber.Map{status=string,message=string,data=[]models.Flight}
// @Failure 400 {object} fiber.Map{status=string,message=string}
// @Router /flights [get]
func (h *FlightController) GetMany(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	flights, err := h.repository.GetMany(ctxWithTimeout)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    flights,
	})
}

// Lấy thông tin chuyến bay theo ID
// @Summary Lấy thông tin chuyến bay
// @Description Lấy thông tin của một chuyến bay dựa trên ID
// @Tags Chuyến bay
// @Accept  json
// @Produce  json
// @Param flightId path int true "ID chuyến bay"
// @Success 200 {object} fiber.Map{status=string,message=string,data=models.Flight}
// @Failure 400 {object} fiber.Map{status=string,message=string}
// @Router /flights/{flightId} [get]
func (h *FlightController) GetOne(ctx *fiber.Ctx) error {
	flightId, _ := strconv.Atoi(ctx.Params("flightId"))

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	flight, err := h.repository.GetOne(ctxWithTimeout, uint(flightId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    flight,
	})
}

// Tạo chuyến bay mới
// @Summary Tạo chuyến bay
// @Description Tạo một chuyến bay mới trong hệ thống
// @Tags Chuyến bay
// @Accept  json
// @Produce  json
// @Param flight body models.Flight true "Dữ liệu chuyến bay"
// @Success 201 {object} fiber.Map{status=string,message=string,data=models.Flight}
// @Failure 400 {object} fiber.Map{status=string,message=string}
// @Router /flights [post]
func (h *FlightController) CreateOne(ctx *fiber.Ctx) error {
	flight := &models.Flight{}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(flight); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	flightAfterCreate, err := h.repository.CreateOne(ctxWithTimeout, flight)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Tạo chuyến bay thành công",
		"data":    flightAfterCreate,
	})
}

// Cập nhật thông tin chuyến bay
// @Summary Cập nhật chuyến bay
// @Description Cập nhật thông tin chuyến bay dựa trên ID
// @Tags Chuyến bay
// @Accept  json
// @Produce  json
// @Param flightId path int true "ID chuyến bay"
// @Param updateData body map[string]interface{} true "Dữ liệu cập nhật"
// @Success 201 {object} fiber.Map{status=string,message=string,data=models.Flight}
// @Failure 400 {object} fiber.Map{status=string,message=string}
// @Router /flights/{flightId} [put]
func (h *FlightController) UpdateOne(ctx *fiber.Ctx) error {
	flightId, _ := strconv.Atoi(ctx.Params("flightId"))
	updateData := make(map[string]interface{})

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	flightAfterUpdate, err := h.repository.UpdateOne(ctxWithTimeout, uint(flightId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Cập nhật chuyến bay thành công",
		"data":    flightAfterUpdate,
	})
}

// Xóa chuyến bay
// @Summary Xóa chuyến bay
// @Description Xóa chuyến bay dựa trên ID
// @Tags Chuyến bay
// @Param flightId path int true "ID chuyến bay"
// @Success 204 "Không có nội dung trả về"
// @Failure 400 {object} fiber.Map{status=string,message=string}
// @Router /flights/{flightId} [delete]
func (h *FlightController) DeleteOne(ctx *fiber.Ctx) error {
	flightId, _ := strconv.Atoi(ctx.Params("flightId"))

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(ctxWithTimeout, uint(flightId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// NewFlightController trả về một đối tượng FlightController mới
func NewFlightController(IFlightRepository repositories.IFlightRepository) *FlightController {
	return &FlightController{
		repository: IFlightRepository,
	}
}
