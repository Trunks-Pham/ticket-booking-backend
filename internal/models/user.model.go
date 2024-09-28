package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Customer UserRole = "customer"
)

// User godoc
// @Description Thông tin người dùng
type User struct {
	gorm.Model
	FirstName      string           `json:"firstName" gorm:"varchar(255);not null" example:"John"`             // Tên người dùng
	LastName       string           `json:"lastName" gorm:"varchar(255);not null" example:"Doe"`               // Họ người dùng
	Email          string           `json:"email" gorm:"varchar(255);not null" example:"john.doe@example.com"` // Email người dùng
	Password       string           `json:"password" gorm:"type:varchar(255);not null"`                        // Mật khẩu người dùng
	PhoneNumber    string           `json:"phoneNumber" gorm:"varchar(255);not null" example:"+841234567890"`  // Số điện thoại
	IdentityID     string           `json:"identityId" gorm:"varchar(255);" example:"123456789"`               // Số chứng minh thư
	Passport       string           `json:"passport" gorm:"varchar(256);" example:"P1234567"`                  // Hộ chiếu
	Role           UserRole         `json:"role" gorm:"text;default:customer" example:"customer"`              // Vai trò người dùng
	Status         bool             `json:"status" gorm:"default:true" example:"true"`                         // Trạng thái tài khoản
	BookingHistory []BookingHistory `json:"bookings"`                                                          // Lịch sử đặt vé
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return
}
