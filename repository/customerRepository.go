package repository

import (
	"FeedbackAPI/models"
	"errors"
	"gorm.io/gorm"
)

var ErrCustomerNotFound = errors.New("customer does not exist")

type CustomerRepository interface {
	Create(newCustomer *models.Customer) error
	GetByName(name string) (*models.Customer, error)
	GetByID(id uint) (*models.Customer, error)
	Delete(id uint) error
	Update(id uint, updatedCustomer models.UpdateCustomerRequest) (*models.Customer, error)
	GetFeedbacks(id uint) ([]models.Feedback, error)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	return &customer, nil
}

func (cr *customerRepository) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := cr.db.Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	return &customer, nil
}

func (cr *customerRepository) Delete(id uint) error {
	result := cr.db.Delete(&models.Customer{}, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return ErrCustomerNotFound
	}
	return nil
}

func (cr *customerRepository) Update(id uint, updatedCustomer models.UpdateCustomerRequest) (*models.Customer, error) {
	result := cr.db.Model(&models.Customer{}).Where("id = ?", id).Updates(updatedCustomer)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, ErrCustomerNotFound
	}

	customer, err := cr.GetByID(id)
	return customer, err
}

func (cr *customerRepository) GetFeedbacks(id uint) ([]models.Feedback, error) {
	var customer models.Customer
	if err := cr.db.Preload("Feedbacks").Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	return customer.Feedbacks, nil
}
