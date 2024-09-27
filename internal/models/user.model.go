package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Customer UserRole = "customer"
)

type User struct {
	gorm.Model
	FirstName      string           `json:"firstName" gorm:"varchar(255);not null"`
	LastName       string           `json:"lastName" gorm:"varchar(255);not null"`
	Email          string           `json:"email" gorm:"varchar(255);not null"`
	Password       string           `json:"password" gorm:"type:varchar(255);not null"`
	PhoneNumber    string           `json:"phoneNumber" gorm:"varchar(255);not null"`
	IdentityID     string           `json:"identityId" gorm:"varchar(255);"`
	Passport       string           `json:"passport" gorm:"varchar(256);"`
	Role           UserRole         `json:"role" gorm:"text;default:customer"`
	Status         bool             `json:"status" gorm:"default:true"`
	BookingHistory []BookingHistory `json:"bookings" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return
}
