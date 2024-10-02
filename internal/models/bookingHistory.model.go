package models

import (
	"gorm.io/gorm"
)

// BookingHistory lưu trữ lịch sử đặt vé của người dùng.
// @Description Đối tượng BookingHistory chứa thông tin về các lần đặt vé của người dùng, bao gồm mã người dùng, vé, tổng số tiền, số ghế, cổng vào và trạng thái đặt vé.
// @Param userId body uint true "Mã người dùng"
// @Param user body User true "Thông tin người dùng"
// @Param ticketId body uint true "Mã vé"
// @Param ticket body Ticket true "Thông tin vé"
// @Param totalAmount body float64 true "Tổng số tiền thanh toán"
// @Param seatNumber body string true "Số ghế"
// @Param gate body string true "Cổng vào"
// @Param status body boolean true "Trạng thái của vé (true: Đã đặt, false: Huỷ)"
type BookingHistory struct {
	gorm.Model
	UserID      uint    `json:"userId" gorm:"not null"`                                                           // Mã người dùng
	User        User    `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // Thông tin người dùng
	TicketID    uint    `json:"ticketId" gorm:"not null"`                                                         // Mã vé
	Ticket      Ticket  `json:"ticket" gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE;onDelete:SET NULL;"` // Thông tin vé
	TotalAmount float64 `json:"totalAmount" gorm:"type:decimal(10,2);not null"`                                   // Tổng số tiền thanh toán
	SeatNumber  string  `json:"seatNumber" gorm:"type:varchar(255);not null"`                                     // Số ghế ngồi
	Gate        string  `json:"gate" gorm:"type:varchar(255);not null"`                                           // Cổng vào
	Status      bool    `json:"status" gorm:"type:boolean;default:true"`                                          // Trạng thái đặt vé (true: Đã đặt, false: Huỷ)
}
