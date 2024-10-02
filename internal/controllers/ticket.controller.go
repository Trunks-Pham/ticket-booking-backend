package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type TicketController struct {
	repository repositories.ITicketRepository
}

// GetMany godoc
// @Summary Lấy nhiều vé
// @Description Lấy danh sách vé có thể lọc theo flightId
// @Tags tickets
// @Produce json
// @Param flightId query string false "Flight ID"
// @Success 200 {object} fiber.Map{"status": string, "message": string, "data": []models.Ticket}
// @Failure 400 {object} fiber.Map{"status": string, "message": string}
// @Router /tickets [get]
func (h *TicketController) GetMany(ctx *fiber.Ctx) error {
	flightIdStr := ctx.Query("flightId")
	var flightId *uint

	if flightIdStr != "" {
		id, err := strconv.ParseUint(flightIdStr, 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "fail",
				"message": "Invalid Flight ID",
			})
		}
		idUint := uint(id)
		flightId = &idUint
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(ctxWithTimeout, flightId)

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

// GetOne godoc
// @Summary Lấy chi tiết vé
// @Description Lấy chi tiết của một vé dựa trên ticketId
// @Tags tickets
// @Produce json
// @Param ticketId path int true "Ticket ID"
// @Success 200 {object} fiber.Map{"status": string, "message": string, "data": fiber.Map{"ticket": models.Ticket}}
// @Failure 400 {object} fiber.Map{"status": string, "message": string}
// @Router /tickets/{ticketId} [get]
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

// CreateOne godoc
// @Summary Tạo mới vé
// @Description Tạo một vé mới
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body models.Ticket true "Thông tin vé"
// @Success 201 {object} fiber.Map{"status": string, "message": string, "data": models.Ticket}
// @Failure 422 {object} fiber.Map{"status": string, "message": string}
// @Router /tickets [post]
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

	ticketResponse, err := h.repository.CreateOne(ctxWithTimeout, ticket)

	if err != nil {
		fmt.Printf("Error creating ticket: %v\n", err)
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

// UpdateOne godoc
// @Summary Cập nhật vé
// @Description Cập nhật thông tin của một vé dựa trên ticketId
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticketId path int true "Ticket ID"
// @Param updateData body map[string]interface{} true "Thông tin cần cập nhật"
// @Success 200 {object} fiber.Map{"status": string, "message": string, "data": models.Ticket}
// @Failure 422 {object} fiber.Map{"status": string, "message": string}
// @Router /tickets/{ticketId} [put]
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

// DeleteOne godoc
// @Summary Xóa vé
// @Description Xóa một vé dựa trên ticketId
// @Tags tickets
// @Param ticketId path int true "Ticket ID"
// @Success 204
// @Failure 400 {object} fiber.Map{"status": string, "message": string}
// @Router /tickets/{ticketId} [delete]
func (h *TicketController) DeleteOne(ctx *fiber.Ctx) error {
	ticketId, _ := strconv.Atoi(ctx.Params("ticketId"))

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(ctxWithTimeout, uint(ticketId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewTicketController(ITicketRepository repositories.ITicketRepository) *TicketController {
	return &TicketController{
		repository: ITicketRepository,
	}
}
