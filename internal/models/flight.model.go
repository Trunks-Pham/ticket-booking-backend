package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Flight godoc
// @Description Thông tin chuyến bay
type Flight struct {
	gorm.Model
	FlightNumber     string    `json:"flightNumber" example:"VN123"`                 // Số hiệu chuyến bay
	DepartureAirport string    `json:"departureAirport" example:"SGN"`               // Sân bay đi
	ArrivalAirport   string    `json:"arrivalAirport" example:"HAN"`                 // Sân bay đến
	DepartureTime    time.Time `json:"departureTime" example:"2024-09-28T10:00:00Z"` // Thời gian cất cánh
	ArrivalTime      time.Time `json:"arrivalTime" example:"2024-09-28T12:00:00Z"`   // Thời gian hạ cánh
	AircraftType     string    `json:"aircraftType" example:"Airbus A320"`           // Loại máy bay
	Status           bool      `json:"status" example:"true"`                        // Trạng thái chuyến bay
	Ticket           []Ticket  `json:"ticket"`                                       // Vé chuyến bay
}

type IFlightRepository interface {
	GetMany(ctx context.Context) ([]*Flight, error)
	GetOne(ctx context.Context, eventId uint) (*Flight, error)
	CreateOne(ctx context.Context, event *Flight) (*Flight, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Flight, error)
	DeleteOne(ctx context.Context, eventId uint) error
}
