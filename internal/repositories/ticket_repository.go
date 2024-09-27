package repositories

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
)

type ITicketRepository interface {
	GetMany(ctx context.Context, eventId *uint) ([]*models.Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error)
	CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error)
	DeleteOne(ctx context.Context, ticketId uint) error
}
