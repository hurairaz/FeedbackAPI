package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Rating     float64
	Comment    string
	CustomerID uint
}

type CreateFeedbackRequest struct {
	Rating  float64 `json:"rating" binding:"required"`
	Comment string  `json:"comment" binding:"required"`
}

type CreateFeedbackResponse struct {
	Success bool
	Error   string
	ID      uint
	Rating  float64
	Comment string
}

type UpdateFeedbackRequest struct {
	Rating  float64 `json:"rating,omitempty"`
	Comment string  `json:"comment,omitempty"`
}

type UpdateFeedbackResponse struct {
	Success  bool
	Error    string
	Feedback Feedback
}
