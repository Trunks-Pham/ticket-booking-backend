package repositories

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IBookingHistoryRepository interface {
	GetMany(ctx context.Context, userId *uint) ([]*models.BookingHistory, error)
	GetOne(ctx context.Context, bookingHistoryId uint) (*models.BookingHistory, error)
	CreateOne(ctx context.Context, bookingHistory *models.BookingHistory) (*models.BookingHistory, error)
	UpdateOne(ctx context.Context, bookingHistoryId uint, updateData map[string]interface{}) (*models.BookingHistory, error)
	DeleteOne(ctx context.Context, bookingHistoryId uint) error
	AssignSeat(ctx context.Context, ticketId uint) (string, error)
}
