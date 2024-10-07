package implement

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"gorm.io/gorm"
	"strconv"
)

type BookingHistoryRepository struct {
	db *gorm.DB
}

func (b BookingHistoryRepository) AssignSeat(ctx context.Context, ticketId uint) (string, error) {
	var seatNumbers []string

	res := b.db.Model(&models.BookingHistory{}).
		Where("ticket_id = ?", ticketId).
		Pluck("seat_numbers", &seatNumbers).Error

	if res.Error != nil {
		// Không làm gì hớt
	}

	maxSeatNumber := 0
	for _, seat := range seatNumbers {
		seatNum, err := strconv.Atoi(seat)
		if err != nil && seatNum > maxSeatNumber {
			maxSeatNumber = seatNum
		}
	}

	newSeatNumber := maxSeatNumber + 1

	return strconv.Itoa(newSeatNumber), nil
}

func (b BookingHistoryRepository) GetMany(ctx context.Context, userId *uint) ([]*models.BookingHistory, error) {
	bookingHistories := []*models.BookingHistory{}

	query := b.db.Model(&models.BookingHistory{})

	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}

	res := query.Preload("User").Preload("Ticket.Flight").Find(&bookingHistories)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookingHistories, nil
}

func (b BookingHistoryRepository) GetOne(ctx context.Context, bookingHistoryId uint) (*models.BookingHistory, error) {
	bookingHistory := &models.BookingHistory{}

	res := b.db.Model(&models.BookingHistory{}).Where("id = ?", bookingHistoryId).Preload("User").Preload("Ticket.Flight").First(bookingHistory)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookingHistory, nil
}

func (b BookingHistoryRepository) CreateOne(ctx context.Context, bookingHistory *models.BookingHistory) (*models.BookingHistory, error) {
	res := b.db.Model(bookingHistory).Create(bookingHistory)

	if res.Error != nil {
		return nil, res.Error
	}

	return bookingHistory, nil
}

func (b BookingHistoryRepository) UpdateOne(ctx context.Context, bookingHistoryId uint, updateData map[string]interface{}) (*models.BookingHistory, error) {
	bookingHistory := &models.BookingHistory{}

	updateRes := b.db.Model(bookingHistory).Where("id = ?", bookingHistoryId).Preload("User").Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return b.GetOne(ctx, bookingHistoryId)
}

func (b BookingHistoryRepository) DeleteOne(ctx context.Context, bookingHistoryId uint) error {
	res := b.db.Delete(&models.BookingHistory{}, bookingHistoryId)
	return res.Error
}

func NewBookingHistoryRepository() repositories.IBookingHistoryRepository {
	return &BookingHistoryRepository{
		db: global.Pdb,
	}
}
