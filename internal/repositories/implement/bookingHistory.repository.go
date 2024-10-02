package implement

import (
	"context"

	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"gorm.io/gorm"
)

type BookingHistoryRepository struct {
	db *gorm.DB
}

// Lấy danh sách lịch sử đặt vé của người dùng
func (r *BookingHistoryRepository) GetUserBookingHistory(ctx context.Context, userId uint) ([]*models.BookingHistory, error) {
	var bookingHistories []*models.BookingHistory

	res := r.db.Model(&models.BookingHistory{}).
		Where("user_id = ?", userId).
		Preload("Ticket.Flight"). // Tải dữ liệu chuyến bay liên quan
		Find(&bookingHistories)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookingHistories, nil
}

// Lưu trữ lịch sử đặt vé
func (r *BookingHistoryRepository) CreateBooking(ctx context.Context, booking *models.BookingHistory) (*models.BookingHistory, error) {
	res := r.db.Model(booking).Create(booking)
	if res.Error != nil {
		return nil, res.Error
	}

	return booking, nil
}

// Tạo một BookingHistoryRepository mới
func NewBookingHistoryRepository() repositories.IBookingHistoryRepository {
	return &BookingHistoryRepository{
		db: global.Pdb,
	}
}
