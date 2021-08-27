package repositories

import (
	"context"
	"go-booking/grpc/customer-grpc/models"
	"go-booking/grpc/customer-grpc/requests"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerByID(context context.Context, id uuid.UUID) (*models.Customer, error)
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=customer port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetCustomerByID(context context.Context, id uuid.UUID) (*models.Customer, error) {
	Customer := models.Customer{}
	if err := m.Where(&models.Customer{ID: id}).First(&Customer).Error; err != nil {
		return nil, err
	}

	return &Customer, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Where(&models.Customer{ID: model.ID}).Updates(&models.Customer{Name: model.Name}).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (m *dbmanager) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	customer, err := m.GetCustomerByID(ctx, id)
	if err != nil {
		return err
	}

	return m.Unscoped().Delete(&customer).Error
}
func (m *dbmanager) ListCustomers(ctx context.Context, req *requests.ListCustomerRequest) ([]*models.Customer, error) {
	customers := []*models.Customer{}

	return customers, nil
}
