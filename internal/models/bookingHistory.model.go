package models

import (
	"gorm.io/gorm"
)

// BookingHistory godoc
// @Description Lịch sử đặt vé của người dùng
type BookingHistory struct {
	gorm.Model
	UserID      uint    `json:"userId" example:"1"`          // ID người dùng
	User        User    `json:"user"`                        // Thông tin người dùng
	TicketID    uint    `json:"ticketId" example:"1"`        // ID vé
	Ticket      Ticket  `json:"ticket"`                      // Thông tin vé
	TotalAmount float64 `json:"totalAmount" example:"99.99"` // Tổng số tiền
	SeatNumber  string  `json:"seatNumber" example:"12A"`    // Số ghế ngồi
	Gate        string  `json:"gate" example:"Gate A"`       // Cổng lên máy bay
	Status      bool    `json:"status" example:"true"`       // Trạng thái đặt vé
}
