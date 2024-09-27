package repositories

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type IAuthRepository interface {
	RegisterUser(ctx context.Context, registerData *models.RegisterCredentials) (*models.User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error)
}
