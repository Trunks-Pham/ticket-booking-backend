package controllers

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FlightController struct {
	repository repositories.IFlightRepository
}

func (h *FlightController) GetMany(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	flight, err := h.repository.GetMany(ctxWithTimeout)

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
		"message": "Flight created",
		"data":    flightAfterCreate,
	})
}

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
		"message": "Flight updated",
		"data":    flightAfterUpdate,
	})
}

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

func NewFlightController(IFlightRepository repositories.IFlightRepository) *FlightController {
	return &FlightController{
		repository: IFlightRepository,
	}
}
