package main_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Test for application initialization
func TestAppInitialization(t *testing.T) {
	app := fiber.New()

	// Kiểm tra xem ứng dụng có khởi tạo thành công hay không
	assert.NotNil(t, app)
}

// Test Swagger Route
func TestSwaggerRoute(t *testing.T) {
	app := fiber.New()
	app.Get("/swagger/*", func(c *fiber.Ctx) error {
		return c.SendString("Swagger Docs")
	})

	// Tạo một yêu cầu HTTP GET tới "/swagger/index.html"
	req := httptest.NewRequest("GET", "/swagger/index.html", nil)
	resp, err := app.Test(req)

	// Kiểm tra nếu có lỗi trong yêu cầu
	assert.NoError(t, err)

	// Kiểm tra mã trạng thái của phản hồi
	assert.Equal(t, 200, resp.StatusCode)

	// Đọc và kiểm tra nội dung của phản hồi
	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)
	assert.Equal(t, "Swagger Docs", string(body))
}
