package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// RegisterCredentials godoc
// @Description Thông tin đăng ký tài khoản người dùng
type RegisterCredentials struct {
	Email       string `json:"email" validate:"required" example:"user@example.com"`    // Email người dùng
	Password    string `json:"password" validate:"required" example:"password123"`      // Mật khẩu người dùng
	FirstName   string `json:"firstName" validate:"required" example:"John"`            // Tên của người dùng
	LastName    string `json:"lastName" validate:"required" example:"Doe"`              // Họ của người dùng
	PhoneNumber string `json:"phoneNumber" validate:"required" example:"+841234567890"` // Số điện thoại
	IdentityID  string `json:"identityId" validate:"omitempty" example:"123456789"`     // Chứng minh thư
	Passport    string `json:"passport" validate:"omitempty" example:"P1234567"`        // Hộ chiếu
}

type LoginCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type IAuthRepository interface {
	RegisterUser(ctx context.Context, registerData *RegisterCredentials) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type IAuthService interface {
	Login(ctx context.Context, loginData *LoginCredentials) (string, *User, error)
	Register(ctx context.Context, registerData *RegisterCredentials) (string, *User, error)
}

// Check if a password matches a hash
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Checks if an email is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
