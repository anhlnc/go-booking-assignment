package repositories

import (
	"context"
	"go-booking/grpc/flight-grpc/models"
	"go-booking/grpc/flight-grpc/requests"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Embeded struct

type FlightRepository interface {
	GetFlightID(context context.Context, id uuid.UUID) (*models.Flight, error)
	FindFlight(context context.Context, req *requests.FindFlightRequest) (*models.Flight, error)
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	ListFlights(ctx context.Context, req *requests.ListFlightRequest) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=flight port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) GetFlightID(ctx context.Context, id uuid.UUID) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{ID: id}).First(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func (m *dbmanager) FindFlight(context context.Context, req *requests.FindFlightRequest) (*models.Flight, error) {
	Flight := models.Flight{}
	if req.Name != "" {
		if err := m.Where(&models.Flight{Name: req.Name}).First(&Flight).Error; err != nil {
			return nil, err
		}
	}
	// if req.From != "" {
	// 	if err := m.Where(&models.Flight{From: req.From}).Find(&Flight).Error; err != nil {
	// 		return nil, err
	// 	}
	// }
	// if req.To != "" {
	// 	if err := m.Where(&models.Flight{To: req.To}).Find(&Flight).Error; err != nil {
	// 		return nil, err
	// 	}
	// }
	// if req.Date != "" {
	// 	if err := m.Where(&models.Flight{Date: req.Date}).Find(&Flight).Error; err != nil {
	// 		return nil, err
	// 	}
	// }
	return &Flight, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Where(&models.Flight{ID: model.ID}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) ListFlights(ctx context.Context, req *requests.ListFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}

	return flights, nil
}
