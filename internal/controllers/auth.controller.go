package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Tạo một instance validator mới để kiểm tra dữ liệu
var validate = validator.New()

// AuthController quản lý các yêu cầu liên quan đến xác thực người dùng
type AuthController struct {
	service services.IAuthService
}

// Login xử lý đăng nhập người dùng
// @Summary      Đăng nhập người dùng
// @Description  Xác thực người dùng với thông tin đăng nhập
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  models.LoginCredentials  true  "Thông tin đăng nhập"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /auth/login [post]
func (h *AuthController) Login(ctx *fiber.Ctx) error {
	// Khởi tạo thông tin đăng nhập từ request body
	loginCreds := &models.LoginCredentials{}

	// Tạo context với thời gian chờ 5 giây
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Phân tích dữ liệu từ request body thành loginCreds
	if err := ctx.BodyParser(&loginCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": err.Error(),
		})
	}

	// Kiểm tra tính hợp lệ của thông tin đăng nhập
	if err := validate.Struct(loginCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": err.Error(),
		})
	}

	// Thực hiện đăng nhập qua service
	token, user, err := h.service.Login(ctxWithTimeout, loginCreds)

	// Xử lý lỗi đăng nhập
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": err.Error(),
		})
	}

	// Trả về phản hồi thành công với token và thông tin người dùng
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "thành công",
		"message": "Đăng nhập thành công",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

// Register xử lý đăng ký người dùng mới
// @Summary      Đăng ký người dùng
// @Description  Đăng ký người dùng mới
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  models.RegisterCredentials  true  "Thông tin đăng ký"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /auth/register [post]
func (h *AuthController) Register(ctx *fiber.Ctx) error {
	// Khởi tạo thông tin đăng ký từ request body
	registerCreds := &models.RegisterCredentials{}

	// Tạo context với thời gian chờ 5 giây
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Phân tích dữ liệu từ request body thành registerCreds
	if err := ctx.BodyParser(&registerCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": err.Error(),
		})
	}

	// Kiểm tra tính hợp lệ của thông tin đăng ký
	if err := validate.Struct(registerCreds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": fmt.Errorf("vui lòng cung cấp thông tin hợp lệ").Error(),
		})
	}

	// Thực hiện đăng ký qua service
	token, user, err := h.service.Register(ctxWithTimeout, registerCreds)

	// Xử lý lỗi đăng ký
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "thất bại",
			"message": err.Error(),
		})
	}

	// Trả về phản hồi thành công với token và thông tin người dùng
	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "thành công",
		"message": "Đăng ký thành công",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

// NewAuthController tạo một instance mới của AuthController
func NewAuthController(service services.IAuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}
