package handlers

import (
	"context"
	"database/sql"
	"go-booking/grpc/flight-grpc/models"
	"go-booking/grpc/flight-grpc/repositories"
	"go-booking/grpc/flight-grpc/requests"
	"go-booking/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightHandler struct {
	pb.UnimplementedBookingFlightServer
	flightRepository repositories.FlightRepository
}

func NewFlightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
	}, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	pRequest := &models.Flight{
		ID:             uuid.New(),
		Name:           in.Name,
		From:           in.From,
		To:             in.To,
		Date:           in.Date,
		Status:         in.Status,
		Available_slot: int(in.AvailableSlot),
	}
	flight, err := h.flightRepository.CreateFlight(ctx, pRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.Flight{
		Id:            flight.ID.String(),
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Date:          flight.Date,
		Status:        flight.Status,
		AvailableSlot: int64(flight.Available_slot),
	}

	return pResponse, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.flightRepository.GetFlightID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Date != "" {
		flight.Date = in.Date
	}
	if in.From != "" {
		flight.From = in.From
	}
	if in.To != "" {
		flight.To = in.To
	}
	if in.Status != "" {
		flight.Status = in.Status
	}
	if in.Name != "" {
		flight.Name = in.Name
	}
	if in.AvailableSlot != 0 {
		flight.Available_slot = int(in.AvailableSlot)
	}

	newFlight, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Flight{
		Id:            newFlight.ID.String(),
		Name:          newFlight.Name,
		From:          newFlight.From,
		To:            newFlight.To,
		Date:          newFlight.Date,
		Status:        newFlight.Status,
		AvailableSlot: int64(newFlight.Available_slot),
	}

	return pResponse, nil
}
func (h *FlightHandler) FindFlight(ctx context.Context, in *pb.FindFlightRequest) (*pb.Flight, error) {
	var (
		flights = &models.Flight{}
		err     error
	)

	flights, err = h.flightRepository.FindFlight(ctx, &requests.FindFlightRequest{
		Name: in.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.Flight{
		Id:            flights.ID.String(),
		Name:          flights.Name,
		From:          flights.From,
		To:            flights.To,
		Date:          flights.Date,
		Status:        flights.Status,
		AvailableSlot: int64(flights.Available_slot),
	}
	return pRes, nil
}

func (h *FlightHandler) ListFlights(ctx context.Context, in *pb.ListFlightRequest) (*pb.ListFlightResponse, error) {
	flights, err := h.flightRepository.ListFlights(ctx, &requests.ListFlightRequest{
		Name: in.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.ListFlightResponse{
		Flights: []*pb.Flight{},
	}

	err = copier.CopyWithOption(&pRes.Flights, &flights, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *FlightHandler) GetFlightID(ctx context.Context, in *pb.GetFlightIDRequest) (*pb.Flight, error) {
	var (
		flight = &models.Flight{}
		err    error
	)

	if in.Id != "" {
		flight, err = h.flightRepository.GetFlightID(ctx, uuid.MustParse(in.Id))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	pRes := &pb.Flight{
		Id:            flight.ID.String(),
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Date:          flight.Date,
		Status:        flight.Status,
		AvailableSlot: int64(flight.Available_slot),
	}

	return pRes, nil
}
