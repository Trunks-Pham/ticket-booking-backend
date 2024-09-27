package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func CheckManagerRole() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		role := ctx.Locals("role")
		if role == nil || role != "manager" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "fail",
				"message": "Access denied",
			})
		}

		return ctx.Next()
	}
}
