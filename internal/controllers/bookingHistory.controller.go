package controllers

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

type BookingHistoryController struct {
	bookingService    services.IBookingHistoryService
	bookingRepository repositories.IBookingHistoryRepository
}

func (h *BookingHistoryController) BookTicket(ctx *fiber.Ctx) error {
	bookingHistory := &models.BookingHistory{}

	userId := ctx.Locals("userId").(uint)

	ctxWithTimeout, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	if err := ctx.BodyParser(bookingHistory); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	bookingHistory.UserID = userId

	booking, err := h.bookingService.BookTicket(ctxWithTimeout, bookingHistory)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket booked successfully",
		"data":    booking,
	})
}

func (h *BookingHistoryController) GetMany(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	userId := ctx.Locals("userId").(uint)

	bookingHistories, err := h.bookingRepository.GetMany(ctxWithTimeout, &userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    bookingHistories,
	})
}

func NewBookingHistoryController(
	bookingService services.IBookingHistoryService,
	bookingRepository repositories.IBookingHistoryRepository,
) *BookingHistoryController {
	return &BookingHistoryController{
		bookingService:    bookingService,
		bookingRepository: bookingRepository,
	}
}
