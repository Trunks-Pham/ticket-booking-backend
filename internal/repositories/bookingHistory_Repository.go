package repositories

import (
	"context"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IBookingHistoryRepository interface {
	GetUserBookingHistory(ctx context.Context, userId uint) ([]*models.BookingHistory, error)
	CreateBooking(ctx context.Context, booking *models.BookingHistory) (*models.BookingHistory, error)
}
