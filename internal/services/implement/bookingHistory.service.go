package implement

import (
	"context"
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/repositories"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/services"
	"github.com/Trunks-Pham/ticket-booking-backend/utils"
)

type BookingHistoryService struct {
	bookingHistoryRepository repositories.IBookingHistoryRepository
	ticketRepository         repositories.ITicketRepository
}

func (b BookingHistoryService) BookTicket(ctx context.Context, bookingHistory *models.BookingHistory) (*models.BookingHistory, error) {
	ticket, err := b.ticketRepository.GetOne(ctx, bookingHistory.TicketID)
	if err != nil {
		return nil, err
	}

	if ticket.AvailableSeat <= 0 {
		return nil, fmt.Errorf("No available seats")
	}

	vatRate := 0.1
	totalAmount := ticket.Price * (1 + vatRate)

	bookingHistory.TotalAmount = totalAmount

	seatNumber, seatNumberErr := b.bookingHistoryRepository.AssignSeat(ctx, ticket.ID)
	if seatNumberErr != nil {
		return nil, seatNumberErr
	}
	bookingHistory.SeatNumber = seatNumber

	bookingHistory.Gate = utils.AssignGate()

	createBookingHistory, err := b.bookingHistoryRepository.CreateOne(ctx, bookingHistory)

	if err != nil {
		return nil, err
	}

	ticket.AvailableSeat--

	_, err = b.ticketRepository.UpdateOne(ctx, ticket.ID, map[string]interface{}{
		"available_seat": ticket.AvailableSeat,
	})

	if err != nil {
		return nil, err
	}

	return createBookingHistory, nil
}

func NewBookingService(
	bookingHistoryRepository repositories.IBookingHistoryRepository,
	ticketRepository repositories.ITicketRepository) services.IBookingHistoryService {
	return &BookingHistoryService{
		bookingHistoryRepository: bookingHistoryRepository,
		ticketRepository:         ticketRepository,
	}
}
