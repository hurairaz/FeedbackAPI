package controllers

import (
	"FeedbackAPI/models"
	"FeedbackAPI/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

type FeedbackController interface {
	CreateFeedback(ctx *fiber.Ctx) error
	GetFeedback(ctx *fiber.Ctx) error
	UpdateFeedback(ctx *fiber.Ctx) error
	DeleteFeedback(ctx *fiber.Ctx) error
}

type feedbackController struct {
	feedbackRepo repository.FeedbackRepository
}

func NewFeedbackController(feedbackRepo repository.FeedbackRepository) FeedbackController {
	return &feedbackController{feedbackRepo: feedbackRepo}
}

func (fc *feedbackController) CreateFeedback(ctx *fiber.Ctx) error {
	var req models.CreateFeedbackRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(models.CreateFeedbackResponse{Success: false,
			Error: err.Error()})
	}

	if strings.TrimSpace(req.Comment) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(models.CreateFeedbackResponse{Success: false,
			Error: "comment can not be empty"})
	}
	if req.Rating <= 0.0 || req.Rating > 5.0 {
		return ctx.Status(http.StatusBadRequest).JSON(models.CreateFeedbackResponse{Success: false,
			Error: "rating should be between 0.0 to 5.0"})
	}

	feedback := models.Feedback{Rating: req.Rating, Comment: req.Comment}
	if err := fc.feedbackRepo.Create(&feedback); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.CreateFeedbackResponse{Success: false,
			Error: err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(models.CreateFeedbackResponse{Success: true, ID: feedback.ID,
		Comment: feedback.Comment, Rating: feedback.Rating})
}

func (fc *feedbackController) GetFeedback(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.GetFeedbackResponse{Success: false,
			Error: "invalid feedback id"})
	}

	feedback, err := fc.feedbackRepo.GetByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.GetFeedbackResponse{Success: false,
			Error: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(models.GetFeedbackResponse{Success: true, Feedback: *feedback})
}

func (fc *feedbackController) UpdateFeedback(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.UpdateFeedbackResponse{Success: false,
			Error: "invalid feedback id"})
	}

	var req models.UpdateFeedbackRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(models.UpdateFeedbackResponse{Success: false,
			Error: err.Error()})
	}

	updatedFeedback, err := fc.feedbackRepo.Update(uint(id), req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.UpdateFeedbackResponse{Success: false,
			Error: err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(models.UpdateFeedbackResponse{Success: true, Feedback: *updatedFeedback})
}

func (fc *feedbackController) DeleteFeedback(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.DeleteFeedbackResponse{Success: false,
			Error: "invalid feedback id"})
	}

	if err := fc.feedbackRepo.Delete(uint(id)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(models.DeleteFeedbackResponse{Success: false,
			Error: err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(models.DeleteFeedbackResponse{Success: true})
}
