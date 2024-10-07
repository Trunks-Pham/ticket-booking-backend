package controllers

import (
	"bytes"
	"io/ioutil"
	"testing"

	"net/http/httptest"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test case for Login
func TestAuthController_Login(t *testing.T) {
	// Setup
	mockService := new(mocks.IAuthService)
	controller := NewAuthController(mockService)

	app := fiber.New()
	app.Post("/auth/login", controller.Login)

	// Test data
	loginCreds := &models.LoginCredentials{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Create a new request
	req := httptest.NewRequest("POST", "/auth/login", fiber.New())
	req.Header.Set("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"email": "test@example.com",
		"password": "password123"
	}`))

	// Mock the service behavior
	mockService.On("Login", mock.Anything, loginCreds).Return("mocked-token", &models.User{ID: "1"}, nil)

	// Perform the request
	resp, err := app.Test(req, 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

// Test case for Register
func TestAuthController_Register(t *testing.T) {
	// Setup
	mockService := new(mocks.IAuthService)
	controller := NewAuthController(mockService)

	app := fiber.New()
	app.Post("/auth/register", controller.Register)

	// Test data
	registerCreds := &models.RegisterCredentials{
		Email:    "test@example.com",
		Password: "password123",
		FullName: "Test User",
	}

	// Create a new request
	req := httptest.NewRequest("POST", "/auth/register", fiber.New())
	req.Header.Set("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"email": "test@example.com",
		"password": "password123",
		"full_name": "Test User"
	}`))

	// Mock the service behavior
	mockService.On("Register", mock.Anything, registerCreds).Return("mocked-token", &models.User{ID: "1"}, nil)

	// Perform the request
	resp, err := app.Test(req, 1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
