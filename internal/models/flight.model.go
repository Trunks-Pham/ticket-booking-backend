package models

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	gorm.Model
	FlightNumber     string    `json:"flightNumber" gorm:"type:varchar(255);not null"`
	DepartureAirport string    `json:"departureAirport" gorm:"type:varchar(255);not null"`
	ArrivalAirport   string    `json:"arrivalAirport" gorm:"type:varchar(255);not null"`
	DepartureTime    time.Time `json:"departureTime" gorm:"type:timestamptz;not null"`
	ArrivalTime      time.Time `json:"arrivalTime" gorm:"type:timestamptz;not null"`
	AircraftType     string    `json:"aircraftType" gorm:"type:varchar(255);not null"`
	Status           bool      `json:"status" gorm:"type:boolean;default:true"`
	Ticket           []Ticket  `json:"ticket" gorm:"foreignKey:FlightID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type FlightUpdate struct {
	FlightNumber     *string    `json:"flightNumber,omitempty"`
	DepartureAirport *string    `json:"departureAirport,omitempty"`
	ArrivalAirport   *string    `json:"arrivalAirport,omitempty"`
	DepartureTime    *time.Time `json:"departureTime,omitempty"`
	ArrivalTime      *time.Time `json:"arrivalTime,omitempty"`
	AircraftType     *string    `json:"aircraftType,omitempty"`
	Status           *bool      `json:"status,omitempty"`
}
