package services

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IAuthService interface {
	Login(ctx context.Context, loginData *models.LoginCredentials) (string, *models.User, error)
	Register(ctx context.Context, registerData *models.RegisterCredentials) (string, *models.User, error)
}
