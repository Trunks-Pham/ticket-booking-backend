package models

import (
	"context"

	"gorm.io/gorm"
)

// Ticket godoc
// @Description Thông tin vé máy bay
type Ticket struct {
	gorm.Model
	FlightID      uint             `json:"flightId" example:"1"`                    // ID chuyến bay
	Flight        Flight           `json:"flight"`                                  // Thông tin chuyến bay
	Price         float64          `json:"price" example:"100.00"`                  // Giá vé
	TicketType    string           `json:"ticketType" example:"Economy"`            // Loại vé (Economy, Business, v.v.)
	Description   string           `json:"description" example:"Vé hạng phổ thông"` // Mô tả vé
	AvailableSeat int              `json:"availableSeat" example:"50"`              // Số ghế còn trống
	Status        bool             `json:"status" example:"true"`                   // Trạng thái vé
	Bookings      []BookingHistory `json:"bookings"`                                // Các lịch sử đặt vé liên quan
}

type ITicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}
