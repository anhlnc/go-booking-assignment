package handlers

import (
	"context"
	"database/sql"
	"go-booking/grpc/booking-grpc/models"
	"go-booking/grpc/booking-grpc/repositories"
	"go-booking/grpc/booking-grpc/requests"
	"go-booking/pb"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingHandler struct {
	customerClient pb.BookingCustomerClient
	flightClient   pb.BookingFlightClient
	pb.UnimplementedBbookingServer
	bookingRepository repositories.BookingRepository
}

func NewBookingHandler(customerClient pb.BookingCustomerClient,
	flightClient pb.BookingFlightClient,
	bookingRepository repositories.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		customerClient:    customerClient,
		flightClient:      flightClient,
		bookingRepository: bookingRepository,
	}, nil
}

func (h *BookingHandler) AssignBooking(ctx context.Context, in *pb.AssignBookingRequest) (*pb.Booking, error) {
	if in.CustomerId == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_id is required")
	}

	if in.FlightId == "" {
		return nil, status.Error(codes.InvalidArgument, "flight_id is required")
	}

	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindCustomerRequest{
		Id: in.CustomerId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer_id is not found")
			}
		} else {
			return nil, err
		}
	}

	flight, err := h.flightClient.GetFlightID(ctx, &pb.GetFlightIDRequest{
		Id: in.FlightId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "flight_id is not found")
			}
		} else {
			return nil, err
		}
	}

	booking := &models.Booking{

		ID:          uuid.New(),
		Customer_id: uuid.MustParse(customer.Id),
		Flight_id:   uuid.MustParse(flight.Id),
		Code:        in.Code,
		Status:      in.Status,
		Booked_date: time.Now().String(),
	}

	res, err := h.bookingRepository.AssignBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Booking{
		Id:         res.ID.String(),
		CustomerId: res.Customer_id.String(),
		FlightId:   res.Flight_id.String(),
		Code:       res.Code,
		Status:     res.Status,
		BookDate:   res.Booked_date,
	}

	return pRes, nil
}

func (h *BookingHandler) ViewBookings(ctx context.Context, in *pb.ViewBookingRequest) (*pb.ViewBookingReponse, error) {
	if in.CustomerId == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_id is required")
	}

	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "code is required")
	}

	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindCustomerRequest{
		Id: in.CustomerId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer_id is not found")
			}
		} else {
			return nil, err
		}
	}
	booking, _ := h.bookingRepository.ViewBookings(ctx, &requests.ViewBookingRequest{
		Customer_id: in.CustomerId,
		Code:        in.Code,
	})
	flight, _ := h.flightClient.GetFlightID(ctx, &pb.GetFlightIDRequest{
		Id: booking.Flight_id.String(),
	})

	pRes := &pb.ViewBookingReponse{
		Code:     booking.Code,
		BookDate: booking.Booked_date,
		Customer: customer,
		Flight:   flight,
	}
	return pRes, nil
}
func (h *BookingHandler) DeleteBooking(ctx context.Context, in *pb.DeleteBookingRequest) (*pb.Empty, error) {
	if err := h.bookingRepository.DeleteBooking(ctx, uuid.MustParse(in.Id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *BookingHandler) GetBookingByID(ctx context.Context, in *pb.GetBookingByIDRequest) (*pb.Booking, error) {
	booking, err := h.bookingRepository.GetBookingByID(ctx, uuid.MustParse(in.Id))
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	pRes := pb.Booking{
		Id:         booking.ID.String(),
		CustomerId: booking.Customer_id.String(),
		FlightId:   booking.Flight_id.String(),
		Code:       booking.Code,
		Status:     booking.Status,
		BookDate:   booking.Booked_date,
	}
	return &pRes, nil
}

func (h *BookingHandler) ListBooking(ctx context.Context, in *pb.ListBookingRequest) (*pb.ListBookingResponse, error) {
	bookings, err := h.bookingRepository.ListBooking(ctx, uuid.MustParse(in.CustomerId))
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	pRes := &pb.ListBookingResponse{
		Bookings: []*pb.Booking{},
	}
	for _, v := range bookings {
		dt := &pb.Booking{
			Id:         v.ID.String(),
			CustomerId: v.Customer_id.String(),
			FlightId:   v.Flight_id.String(),
			Code:       v.Code,
			Status:     v.Status,
			BookDate:   v.Booked_date,
		}
		pRes.Bookings = append(pRes.Bookings, dt)
	}

	return pRes, nil
}
