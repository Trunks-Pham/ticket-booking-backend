package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/controllers"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setup() (*controllers.FlightController, *mocks.IFlightRepository) {
	mockRepo := new(mocks.IFlightRepository)
	controller := controllers.NewFlightController(mockRepo)
	return controller, mockRepo
}

// Test GetMany flights
func TestFlightController_GetMany_Success(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Get("/flights", controller.GetMany)

	// Mock dữ liệu chuyến bay
	flights := []models.Flight{
		{ID: 1, Name: "Flight 1"},
		{ID: 2, Name: "Flight 2"},
	}

	mockRepo.On("GetMany", mock.Anything).Return(flights, nil)

	req := httptest.NewRequest(fiber.MethodGet, "/flights", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestFlightController_GetMany_Failure(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Get("/flights", controller.GetMany)

	mockRepo.On("GetMany", mock.Anything).Return(nil, errors.New("database error"))

	req := httptest.NewRequest(fiber.MethodGet, "/flights", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Test GetOne flight
func TestFlightController_GetOne_Success(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Get("/flights/:flightId", controller.GetOne)

	flight := models.Flight{ID: 1, Name: "Flight 1"}

	mockRepo.On("GetOne", mock.Anything, uint(1)).Return(&flight, nil)

	req := httptest.NewRequest(fiber.MethodGet, "/flights/1", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestFlightController_GetOne_Failure(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Get("/flights/:flightId", controller.GetOne)

	mockRepo.On("GetOne", mock.Anything, uint(1)).Return(nil, errors.New("flight not found"))

	req := httptest.NewRequest(fiber.MethodGet, "/flights/1", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Test CreateOne flight
func TestFlightController_CreateOne_Success(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Post("/flights", controller.CreateOne)

	newFlight := models.Flight{Name: "Flight 1"}
	createdFlight := newFlight
	createdFlight.ID = 1

	mockRepo.On("CreateOne", mock.Anything, &newFlight).Return(&createdFlight, nil)

	payload, err := json.Marshal(newFlight)
	require.NoError(t, err)

	req := httptest.NewRequest(fiber.MethodPost, "/flights", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestFlightController_CreateOne_Failure(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Post("/flights", controller.CreateOne)

	newFlight := models.Flight{Name: "Flight 1"}

	mockRepo.On("CreateOne", mock.Anything, &newFlight).Return(nil, errors.New("failed to create flight"))

	payload, err := json.Marshal(newFlight)
	require.NoError(t, err)

	req := httptest.NewRequest(fiber.MethodPost, "/flights", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Test UpdateOne flight
func TestFlightController_UpdateOne_Success(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Put("/flights/:flightId", controller.UpdateOne)

	updateData := map[string]interface{}{
		"Name": "Updated Flight",
	}

	updatedFlight := models.Flight{ID: 1, Name: "Updated Flight"}

	mockRepo.On("UpdateOne", mock.Anything, uint(1), updateData).Return(&updatedFlight, nil)

	payload, err := json.Marshal(updateData)
	require.NoError(t, err)

	req := httptest.NewRequest(fiber.MethodPut, "/flights/1", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestFlightController_UpdateOne_Failure(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Put("/flights/:flightId", controller.UpdateOne)

	updateData := map[string]interface{}{
		"Name": "Updated Flight",
	}

	mockRepo.On("UpdateOne", mock.Anything, uint(1), updateData).Return(nil, errors.New("failed to update flight"))

	payload, err := json.Marshal(updateData)
	require.NoError(t, err)

	req := httptest.NewRequest(fiber.MethodPut, "/flights/1", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

// Test DeleteOne flight
func TestFlightController_DeleteOne_Success(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Delete("/flights/:flightId", controller.DeleteOne)

	mockRepo.On("DeleteOne", mock.Anything, uint(1)).Return(nil)

	req := httptest.NewRequest(fiber.MethodDelete, "/flights/1", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
}

func TestFlightController_DeleteOne_Failure(t *testing.T) {
	controller, mockRepo := setup()
	app := fiber.New()
	app.Delete("/flights/:flightId", controller.DeleteOne)

	mockRepo.On("DeleteOne", mock.Anything, uint(1)).Return(errors.New("failed to delete flight"))

	req := httptest.NewRequest(fiber.MethodDelete, "/flights/1", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
