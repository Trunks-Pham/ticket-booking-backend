package controllers

import (
	"context"
	"net/http"
	"testing"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockTicketRepository struct {
	mock.Mock
}

func (m *MockTicketRepository) GetMany(ctx context.Context, flightId *uint) ([]models.Ticket, error) {
	args := m.Called(ctx, flightId)
	return args.Get(0).([]models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	args := m.Called(ctx, ticketId)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	args := m.Called(ctx, ticket)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	args := m.Called(ctx, ticketId, updateData)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) DeleteOne(ctx context.Context, ticketId uint) error {
	args := m.Called(ctx, ticketId)
	return args.Error(0)
}

// Unit tests

func TestTicketController_GetMany(t *testing.T) {
	// Setup
	mockRepo := new(MockTicketRepository)
	controller := NewTicketController(mockRepo)

	app := fiber.New()

	// Define route
	app.Get("/tickets", controller.GetMany)

	// Mock data
	tickets := []models.Ticket{
		{ID: 1, FlightID: 123, SeatNumber: "A1"},
		{ID: 2, FlightID: 123, SeatNumber: "B2"},
	}
	mockRepo.On("GetMany", mock.Anything, mock.AnythingOfType("*uint")).Return(tickets, nil)

	// Test request
	req := utils.Request{
		Method: http.MethodGet,
		Path:   "/tickets?flightId=123",
	}

	// Test handler
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify repository was called
	mockRepo.AssertExpectations(t)
}

func TestTicketController_GetOne(t *testing.T) {
	// Setup
	mockRepo := new(MockTicketRepository)
	controller := NewTicketController(mockRepo)

	app := fiber.New()
	app.Get("/tickets/:ticketId", controller.GetOne)

	// Mock data
	ticket := &models.Ticket{ID: 1, FlightID: 123, SeatNumber: "A1"}
	mockRepo.On("GetOne", mock.Anything, uint(1)).Return(ticket, nil)

	// Test request
	req := utils.Request{
		Method: http.MethodGet,
		Path:   "/tickets/1",
	}

	// Test handler
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify repository was called
	mockRepo.AssertExpectations(t)
}

func TestTicketController_CreateOne(t *testing.T) {
	// Setup
	mockRepo := new(MockTicketRepository)
	controller := NewTicketController(mockRepo)

	app := fiber.New()
	app.Post("/tickets", controller.CreateOne)

	// Mock data
	ticket := &models.Ticket{ID: 1, FlightID: 123, SeatNumber: "A1"}
	mockRepo.On("CreateOne", mock.Anything, ticket).Return(ticket, nil)

	// Test request
	req := utils.Request{
		Method: http.MethodPost,
		Path:   "/tickets",
		Body:   `{"flightId": 123, "seatNumber": "A1"}`,
	}

	// Test handler
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Verify repository was called
	mockRepo.AssertExpectations(t)
}

func TestTicketController_UpdateOne(t *testing.T) {
	// Setup
	mockRepo := new(MockTicketRepository)
	controller := NewTicketController(mockRepo)

	app := fiber.New()
	app.Put("/tickets/:ticketId", controller.UpdateOne)

	// Mock data
	ticket := &models.Ticket{ID: 1, FlightID: 123, SeatNumber: "A1"}
	mockRepo.On("UpdateOne", mock.Anything, uint(1), mock.AnythingOfType("map[string]interface {}")).Return(ticket, nil)

	// Test request
	req := utils.Request{
		Method: http.MethodPut,
		Path:   "/tickets/1",
		Body:   `{"seatNumber": "A2"}`,
	}

	// Test handler
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify repository was called
	mockRepo.AssertExpectations(t)
}

func TestTicketController_DeleteOne(t *testing.T) {
	// Setup
	mockRepo := new(MockTicketRepository)
	controller := NewTicketController(mockRepo)

	app := fiber.New()
	app.Delete("/tickets/:ticketId", controller.DeleteOne)

	// Mock data
	mockRepo.On("DeleteOne", mock.Anything, uint(1)).Return(nil)

	// Test request
	req := utils.Request{
		Method: http.MethodDelete,
		Path:   "/tickets/1",
	}

	// Test handler
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Verify repository was called
	mockRepo.AssertExpectations(t)
}
