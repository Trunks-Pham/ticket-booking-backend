package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// authRoute thiết lập các route cho xác thực người dùng (auth).
// @param route là Fiber Router được sử dụng để định nghĩa các đường dẫn.
// @param authService là service xử lý các tác vụ xác thực người dùng.
func authRoute(route fiber.Router, authService models.IAuthService) {
	authController := controllers.NewAuthController(authService)

	// @Summary User Login
	// @Description Xác thực người dùng và trả về token JWT
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param login body models.LoginCredentials true "Thông tin đăng nhập"
	// @Success 200 {object} string "JWT Token"
	// @Failure 401 {object} fiber.Map "Unauthorized"
	// @Router /api/auth/login [post]
	route.Post("/login", authController.Login)

	// @Summary Đăng ký người dùng mới
	// @Description Đăng ký tài khoản người dùng mới
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param register body models.RegisterCredentials true "Thông tin đăng ký"
	// @Success 201 {object} models.User
	// @Failure 400 {object} fiber.Map "Bad Request"
	// @Router /api/auth/register [post]
	route.Post("/register", authController.Register)
}
