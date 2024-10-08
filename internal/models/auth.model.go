package models

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type RegisterCredentials struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	IdentityID  string `json:"identityId" validate:"omitempty"`
	Passport    string `json:"passport" validate:"omitempty"`
}

type LoginCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
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
