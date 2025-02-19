package controllers

import (
	"FeedbackAPI/auth"
	"FeedbackAPI/models"
	"FeedbackAPI/repository"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type CustomerController interface {
	SignInCustomer(ctx *fiber.Ctx) error
	SignUpCustomer(ctx *fiber.Ctx) error
	GetCustomer(ctx *fiber.Ctx) error
	UpdateCustomer(ctx *fiber.Ctx) error
	DeleteCustomer(ctx *fiber.Ctx) error
	GetCustomerFeedbacks(ctx *fiber.Ctx) error
}

type customerController struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerController(customerRepo repository.CustomerRepository) CustomerController {
	return &customerController{customerRepo: customerRepo}
}

func (cc *customerController) SignInCustomer(ctx *fiber.Ctx) error {
	var req models.SignInCustomerRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(models.SignInCustomerResponse{Success: false,
			Error: err.Error()})
	}

	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Password) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(models.SignInCustomerResponse{Success: false,
			Error: "name or password can not be empty"})
	}

	customer, err := cc.customerRepo.GetByName(req.Name)
	if err != nil || customer == nil {
		return ctx.Status(http.StatusNotFound).JSON(models.SignInCustomerResponse{Success: false,
			Error: "customer does not exist"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(models.SignInCustomerResponse{Success: false,
			Error: "password is incorrect"})
	}

	token, err := auth.GenerateJWT(customer.ID)
	if err != nil || token == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(models.SignInCustomerResponse{Success: false,
			Error: "could not generate token"})
	}

	return ctx.Status(http.StatusOK).JSON(models.SignInCustomerResponse{Success: true, ID: customer.ID,
		Name: customer.Name, Token: token})

}

func (cc *customerController) SignUpCustomer(ctx *fiber.Ctx) error {
	var req models.SignUpCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(models.SignUpCustomerResponse{Success: false,
			Error: err.Error()})
	}

	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Password) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(models.SignUpCustomerResponse{Success: false,
			Error: "password or name can not be empty"})
	}

	if _, err := cc.customerRepo.GetByName(req.Name); err == nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.SignUpCustomerResponse{Success: false,
			Error: "customer with the same name already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.SignUpCustomerResponse{Success: false,
			Error: err.Error()})
	}
	hashedPass := string(hashedPassword)

	customer := models.Customer{Name: req.Name, Password: hashedPass}

	if err = cc.customerRepo.Create(&customer); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.SignUpCustomerResponse{Success: false,
			Error: err.Error()})
	}

	token, err := auth.GenerateJWT(customer.ID)
	if err != nil || token == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(models.SignInCustomerResponse{Success: false,
			Error: "could not generate token"})
	}

	return ctx.Status(http.StatusCreated).JSON(models.SignUpCustomerResponse{Success: true,
		ID: customer.ID, Name: customer.Name, Token: token})
}

func (cc *customerController) UpdateCustomer(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.UpdateCustomerResponse{Success: false,
			Error: err.Error()})
	}

	customer, err := cc.customerRepo.GetByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusBadRequest)
	}

	var req models.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity)
	}

}

func (cc *customerController) GetCustomer(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.GetCustomerResponse{Success: false,
			Error: "invalid customer id"})
	}

	customer, err := cc.customerRepo.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(models.GetCustomerResponse{Success: false,
				Error: "customer does not exist"})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(models.GetCustomerResponse{Success: false,
			Error: "failed to fetch customer"})
	}

	return ctx.Status(http.StatusOK).JSON(models.GetCustomerResponse{Success: true, Customer: *customer})
}

func (cc *customerController) DeleteCustomer(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.DeleteCustomerResponse{Success: false,
			Error: "invalid customer id"})
	}

	if err := cc.customerRepo.Delete(uint(id)); err != nil {
		if errors.Is(err, repository.ErrCustomerNotFound) {
			return ctx.Status(http.StatusBadRequest).JSON(models.DeleteCustomerResponse{Success: false,
				Error: err.Error()})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(models.DeleteCustomerResponse{Success: false,
			Error: fmt.Sprintf("failed to delete customer with id: %v", id)})
	}
	return ctx.Status(http.StatusOK).JSON(models.DeleteCustomerResponse{Success: true})

}

func (cc *customerController) GetCustomerFeedbacks(ctx *fiber.Ctx) error {
	return nil
}
