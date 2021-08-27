package repositories

import (
	"context"
	"go-booking/grpc/booking-grpc/models"
	"go-booking/grpc/booking-grpc/requests"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Embeded struct

type BookingRepository interface {
	GetBookingByID(context context.Context, id uuid.UUID) (*models.Booking, error)
	AssignBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	DeleteBooking(ctx context.Context, id uuid.UUID) error
	ViewBookings(ctx context.Context, req *requests.ViewBookingRequest) (*models.Booking, error)
	ListBooking(ctx context.Context, id uuid.UUID) ([]*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=booking port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetBookingByID(context context.Context, id uuid.UUID) (*models.Booking, error) {
	Booking := models.Booking{}
	if err := m.Where(&models.Booking{ID: id}).First(&Booking).Error; err != nil {
		return nil, err
	}

	return &Booking, nil
}

func (m *dbmanager) AssignBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) DeleteBooking(ctx context.Context, id uuid.UUID) error {
	booking, err := m.GetBookingByID(ctx, id)
	if err != nil {
		return err
	}

	return m.Unscoped().Delete(&booking).Error
}
func (m *dbmanager) ViewBookings(ctx context.Context, req *requests.ViewBookingRequest) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Code: req.Code, Customer_id: uuid.MustParse(req.Customer_id)}).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (m *dbmanager) ListBooking(ctx context.Context, id uuid.UUID) ([]*models.Booking, error) {
	bookings := []*models.Booking{}
	if err := m.Where(&models.Booking{Customer_id: id}).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
