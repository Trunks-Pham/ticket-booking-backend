package models

import (
	"gorm.io/gorm"
)

type BookingHistory struct {
	gorm.Model
	UserID      uint    `json:"userId" gorm:"not null"`
	User        User    `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TicketID    uint    `json:"ticketId" gorm:"not null"`
	Ticket      Ticket  `json:"ticket" gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE;onDelete:SET NULL;"`
	TotalAmount float64 `json:"totalAmount" gorm:"type:decimal(10,2);not null"`
	SeatNumber  string  `json:"seatNumber" gorm:"type:varchar(255);not null"`
	Gate        string  `json:"gate" gorm:"type:varchar(255);not null"`
	Status      bool    `json:"status" gorm:"type:boolean;default:true"`
}
