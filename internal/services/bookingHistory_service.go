package services

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IBookingHistoryService interface {
	BookTicket(ctx context.Context, bookingHistory *models.BookingHistory) (*models.BookingHistory, error)
}
