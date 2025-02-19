package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Password  string
	Feedbacks []Feedback `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
}

type SignUpCustomerRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpCustomerResponse struct {
	Success bool
	Error   string
	ID      uint
	Name    string
	Token   string
}

type SignInCustomerRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInCustomerResponse struct {
	Success bool
	Error   string
	ID      uint
	Name    string
	Token   string
}

type DeleteCustomerResponse struct {
	Success bool
	Error   string
}

type GetCustomerResponse struct {
	Success  bool
	Error    string
	Customer Customer
}

type UpdateCustomerRequest map[string]interface{}

type UpdateCustomerResponse struct {
	Success  bool
	Error    string
	Customer Customer
}
