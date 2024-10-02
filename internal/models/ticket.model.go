package models

import (
	"gorm.io/gorm"
)

// Ticket đại diện cho thông tin vé máy bay.
// @Description Đối tượng Ticket chứa các thông tin về vé máy bay như mã chuyến bay, giá vé, loại vé, mô tả, số ghế trống và trạng thái.
// @Param flightId body uint true "ID của chuyến bay liên quan đến vé"
// @Param price body float64 true "Giá vé"
// @Param ticketType body string true "Loại vé (ví dụ: Hạng nhất, Hạng thương gia, Hạng phổ thông)"
// @Param description body string true "Mô tả vé (thông tin chi tiết về vé)"
// @Param availableSeat body int true "Số lượng ghế còn trống"
// @Param status body boolean true "Trạng thái vé (true: Còn khả dụng, false: Đã bán hết)"
type Ticket struct {
	gorm.Model
	FlightID      uint             `json:"flightId" gorm:"not null"`                                                           // ID chuyến bay
	Flight        Flight           `json:"flight" gorm:"foreignKey:FlightID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`    // Thông tin chuyến bay
	Price         float64          `json:"price" gorm:"type:decimal(10, 2);not null"`                                          // Giá vé
	TicketType    string           `json:"ticketType" gorm:"type:varchar(255);not null"`                                       // Loại vé (ví dụ: Hạng nhất, Hạng thương gia)
	Description   string           `json:"description" gorm:"type:text;not null"`                                              // Mô tả vé
	AvailableSeat int              `json:"availableSeat"`                                                                      // Số ghế trống
	Status        bool             `json:"status" gorm:"default:true"`                                                         // Trạng thái vé (true: Còn khả dụng, false: Hết vé)
	Bookings      []BookingHistory `json:"bookings" gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Các lịch sử đặt vé liên quan
}
