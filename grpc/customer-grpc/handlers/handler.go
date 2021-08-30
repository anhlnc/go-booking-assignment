package handlers

import (
	"context"
	"database/sql"
	"go-booking/grpc/customer-grpc/models"
	"go-booking/grpc/customer-grpc/repositories"
	"go-booking/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedBookingCustomerServer
	bookingClient      pb.BbookingClient
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(bookingClient pb.BbookingClient, customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		bookingClient:      bookingClient,
		customerRepository: customerRepository,
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	pRequest := &models.Customer{
		ID:           uuid.New(),
		Name:         in.Name,
		Address:      in.Address,
		License_id:   in.LicenseId,
		Phone_number: in.PhoneNumber,
		Email:        in.Email,
		Password:     in.Password,
		Active:       false,
	}
	customer, err := h.customerRepository.CreateCustomer(ctx, pRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.Customer{
		Id:          customer.ID.String(),
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.License_id,
		PhoneNumber: customer.Phone_number,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      true,
	}

	return pResponse, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Name != "" {
		customer.Name = in.Name
	}
	if in.PhoneNumber != "" {
		customer.Phone_number = in.PhoneNumber
	}
	if in.Address != "" {
		customer.Address = in.Address
	}
	if in.Password != "" {
		customer.Password = in.Password
	}

	if in.Email != "" {
		customer.Email = in.Email
	}
	if in.LicenseId != "" {
		customer.License_id = in.LicenseId
	}

	if in.Active != customer.Active {
		customer.Active = in.Active
	}

	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Customer{
		Id:          newCustomer.ID.String(),
		Name:        newCustomer.Name,
		Address:     newCustomer.Address,
		LicenseId:   newCustomer.License_id,
		PhoneNumber: newCustomer.Phone_number,
		Email:       newCustomer.Email,
		Password:    newCustomer.Password,
		Active:      newCustomer.Active,
	}

	return pResponse, nil
}
func (h *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindCustomerRequest) (*pb.Customer, error) {
	var (
		customer = &models.Customer{}
		err      error
	)

	if in.Id != "" {
		customer, err = h.customerRepository.GetCustomerByID(ctx, uuid.MustParse(in.Id))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	pRes := &pb.Customer{
		Id:          customer.ID.String(),
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.License_id,
		PhoneNumber: customer.Phone_number,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      true,
	}

	return pRes, nil
}

func (h *CustomerHandler) HistoryBooking(ctx context.Context, in *pb.HistoryBookingRequest) (*pb.ListBookedResponse, error) {
	bookeds, err := h.bookingClient.ListBooking(ctx, &pb.ListBookingRequest{
		CustomerId: in.Id,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	pRes := &pb.ListBookedResponse{
		Bookeds: []*pb.Booked{},
	}

	for _, v := range bookeds.Bookings {
		dt := &pb.Booked{
			CustomerId: v.CustomerId,
			FlightId:   v.FlightId,
			Code:       v.Code,
			Status:     v.Status,
			BookDate:   v.BookDate,
		}
		pRes.Bookeds = append(pRes.Bookeds, dt)
	}

	return pRes, nil
}

func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.Customer, error) {
	customer, err := h.customerRepository.GetCustomerByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Password != "" {
		customer.Password = in.Password
	}
	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.Customer{
		Id:          newCustomer.ID.String(),
		Name:        newCustomer.Name,
		Address:     newCustomer.Address,
		LicenseId:   newCustomer.License_id,
		PhoneNumber: newCustomer.Phone_number,
		Email:       newCustomer.Email,
		Password:    newCustomer.Password,
		Active:      newCustomer.Active,
	}

	return pResponse, nil
}
