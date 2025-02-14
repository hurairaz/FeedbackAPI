package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Rating     int
	Comment    string
	CustomerID uint
}
