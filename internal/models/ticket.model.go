package models

import (
	"context"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	FlightID      uint             `json:"flightId" gorm:"not null"`
	Flight        Flight           `json:"flight" gorm:"foreignKey:FlightID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price         float64          `json:"price" gorm:"type:decimal(10, 2);not null"`
	TicketType    string           `json:"ticketType" gorm:"type:varchar(255);not null"`
	Description   string           `json:"description" gorm:"type:text;not null"`
	AvailableSeat int              `json:"availableSeat"`
	Status        bool             `json:"status" gorm:"default:true"`
	Bookings      []BookingHistory `json:"bookings" gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ITicketRepository interface {
	GetMany(ctx context.Context, eventId *uint) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
	DeleteOne(ctx context.Context, ticketId uint) error
}
