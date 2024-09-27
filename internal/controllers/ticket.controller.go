package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

type TicketController struct {
	repository models.ITicketRepository
}

func (h *TicketController) GetMany(ctx *fiber.Ctx) error {
	fmt.Print("cc")

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(ctxWithTimeout)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (h *TicketController) GetOne(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticketId, _ := strconv.Atoi(ctx.Params("ticketId"))

	ticket, err := h.repository.GetOne(ctxWithTimeout, uint(ticketId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data": &fiber.Map{
			"ticket": ticket,
		},
	})
}

func (h *TicketController) CreateOne(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket := &models.Ticket{}

	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	ticketResponse, err := h.repository.CreateOne(ctxWithTimeout)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket created",
		"data":    ticketResponse,
	})
}

func (h *TicketController) UpdateOne(ctx *fiber.Ctx) error {
	ticketId, _ := strconv.Atoi(ctx.Params("ticketId"))
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

	ticketAfterUpdate, err := h.repository.UpdateOne(ctxWithTimeout, uint(ticketId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket updated",
		"data":    ticketAfterUpdate,
	})
}

func NewTicketController(ITicketRepository models.ITicketRepository) *TicketController {
	return &TicketController{
		repository: ITicketRepository,
	}
}
