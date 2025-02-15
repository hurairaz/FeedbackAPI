package repository

import (
	"FeedbackAPI/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(newCustomer *models.Customer) error
	GetByName(name string) (*models.Customer, error)
	GetByID(id uint) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (cr *customerRepository) Create(newCustomer *models.Customer) error {
	return cr.db.Create(newCustomer).Error
}

func (cr *customerRepository) GetByName(name string) (*models.Customer, error) {
	var customer models.Customer
	if err := cr.db.Where("name = ?", name).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *customerRepository) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := cr.db.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
