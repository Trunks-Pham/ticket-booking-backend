package routes

import (
	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func authRoute(route fiber.Router, authService services.IAuthService) {
	authController := controllers.NewAuthController(authService)

	route.Post("/login", authController.Login)
	route.Post("/register", authController.Register)
}
