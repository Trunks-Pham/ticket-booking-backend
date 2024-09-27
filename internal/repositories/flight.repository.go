package repositories

import (
	"context"

	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

func (r *FlightRepository) GetMany(ctx context.Context) ([]*models.Flight, error) {
	flights := []*models.Flight{}

	res := r.db.Model(&models.Flight{}).Find(&flights)

	if res.Error != nil {
		return nil, res.Error
	}

	return flights, nil
}

func (r *FlightRepository) GetOne(ctx context.Context, flightId uint) (*models.Flight, error) {
	event := &models.Flight{}

	res := r.db.Model(event).Where("id = ?", flightId).First(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *FlightRepository) CreateOne(ctx context.Context, flight *models.Flight) (*models.Flight, error) {
	res := r.db.Model(flight).Create(flight)

	if res.Error != nil {
		return nil, res.Error
	}

	return flight, nil
}

func (r *FlightRepository) UpdateOne(ctx context.Context, flightId uint, updateData map[string]interface{}) (*models.Flight, error) {
	flight := &models.Flight{}

	updateRes := r.db.Model(flight).Where("id = ?", flightId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return r.GetOne(ctx, flightId)
}

func (r *FlightRepository) DeleteOne(ctx context.Context, flightId uint) error {
	res := r.db.Delete(&models.Flight{}, flightId)
	return res.Error
}

func NewFlightRepository() models.IFlightRepository {
	return &FlightRepository{
		db: global.Pdb,
	}
}
