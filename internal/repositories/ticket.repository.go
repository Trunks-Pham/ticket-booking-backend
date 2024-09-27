package repositories

import (
	"context"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context, flightId *uint) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	query := r.db.Model(&models.Ticket{})

	if flightId != nil {
		query = query.Where("flight_id = ?", *flightId)
	}

	res := query.Preload("Flight").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", ticketId).Preload("Flight").First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, ticket.ID)
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return r.GetOne(ctx, ticketId)
}

func (r *TicketRepository) DeleteOne(ctx context.Context, ticketId uint) error {
	res := r.db.Delete(&models.Ticket{}, "id = ?", ticketId)
	return res.Error
}

func NewTicketRepository() models.ITicketRepository {
	return &TicketRepository{
		db: global.Pdb,
	}
}
