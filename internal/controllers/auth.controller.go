package controllers

import (
	"context"
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type AuthController struct {
	service services.IAuthService
}

func (h *AuthController) Login(ctx *fiber.Ctx) error {
	loginCreds := &models.LoginCredentials{}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&loginCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(loginCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	token, user, err := h.service.Login(ctxWithTimeout, loginCreds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully logged in",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthController) Register(ctx *fiber.Ctx) error {
	fmt.Printf("cc")
	registerCreds := &models.RegisterCredentials{}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&registerCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(registerCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": fmt.Errorf("please, provide a valid information").Error(),
		})
	}

	token, user, err := h.service.Register(ctxWithTimeout, registerCreds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully registered",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func NewAuthController(service services.IAuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}
