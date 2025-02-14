package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Password  string
	Feedbacks []Feedback `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}
