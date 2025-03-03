package repository

import (
	"FeedbackAPI/models"
	"errors"
	"gorm.io/gorm"
)

var ErrFeedbackNotFound = errors.New("feedback does not exist")

type FeedbackRepository interface {
	Create(newFeedback *models.Feedback) error
	GetByID(id uint) (*models.Feedback, error)
	Delete(id uint) error
	Update(id uint, updatedCustomer models.UpdateCustomerRequest) (*models.Feedback, error)
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (fr *feedbackRepository) Create(newFeedback *models.Feedback) error {
	return fr.db.Create(newFeedback).Error
}

func (fr *feedbackRepository) GetByID(id uint) (*models.Feedback, error) {
	var feedback models.Feedback
	if err := fr.db.Where("id = ?", id).First(&feedback).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFeedbackNotFound
		}
		return nil, err
	}
	return &feedback, nil
}

func (fr *feedbackRepository) Delete(id uint) error {
	result := fr.db.Delete(&models.Feedback{}, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return ErrFeedbackNotFound
	}
	return nil
}

func (fr *feedbackRepository) Update(id uint, updatedCustomer models.UpdateCustomerRequest) (*models.Feedback, error) {
	result := fr.db.Model(&models.Feedback{}).Where("id = ?", id).Updates(updatedCustomer)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return nil, ErrFeedbackNotFound
	}

	feedback, err := fr.GetByID(id)
	return feedback, err
}
