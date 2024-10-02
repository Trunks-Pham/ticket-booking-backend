package models

import (
	"time"

	"gorm.io/gorm"
)

// Flight đại diện cho thông tin chuyến bay.
// @Description Đối tượng Flight chứa các thông tin về chuyến bay như số hiệu chuyến bay, sân bay khởi hành và hạ cánh, thời gian khởi hành và hạ cánh, loại máy bay và trạng thái chuyến bay.
// @Param flightNumber body string true "Số hiệu chuyến bay"
// @Param departureAirport body string true "Sân bay khởi hành"
// @Param arrivalAirport body string true "Sân bay hạ cánh"
// @Param departureTime body string true "Thời gian khởi hành"
// @Param arrivalTime body string true "Thời gian hạ cánh"
// @Param aircraftType body string true "Loại máy bay"
// @Param status body boolean true "Trạng thái chuyến bay (true: Hoạt động, false: Đã huỷ)"
type Flight struct {
	gorm.Model
	FlightNumber     string    `json:"flightNumber" gorm:"type:varchar(255);not null"`                                  // Số hiệu chuyến bay
	DepartureAirport string    `json:"departureAirport" gorm:"type:varchar(255);not null"`                              // Sân bay khởi hành
	ArrivalAirport   string    `json:"arrivalAirport" gorm:"type:varchar(255);not null"`                                // Sân bay hạ cánh
	DepartureTime    time.Time `json:"departureTime" gorm:"type:timestamptz;not null"`                                  // Thời gian khởi hành
	ArrivalTime      time.Time `json:"arrivalTime" gorm:"type:timestamptz;not null"`                                    // Thời gian hạ cánh
	AircraftType     string    `json:"aircraftType" gorm:"type:varchar(255);not null"`                                  // Loại máy bay
	Status           bool      `json:"status" gorm:"type:boolean;default:true"`                                         // Trạng thái chuyến bay (true: Hoạt động, false: Đã huỷ)
	Ticket           []Ticket  `json:"ticket" gorm:"foreignKey:FlightID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Vé liên quan đến chuyến bay
}
