package repositories

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IFlightRepository interface {
	GetMany(ctx context.Context) ([]*models.Flight, error)
	GetOne(ctx context.Context, flightId uint) (*models.Flight, error)
	CreateOne(ctx context.Context, flight *models.Flight) (*models.Flight, error)
	UpdateOne(ctx context.Context, flightId uint, updateData map[string]interface{}) (*models.Flight, error)
	DeleteOne(ctx context.Context, flightId uint) error
}
