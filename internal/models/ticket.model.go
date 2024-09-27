package models

import (
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
