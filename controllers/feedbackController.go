package controllers

import (
	"FeedbackAPI/repository"
	"github.com/gofiber/fiber/v2"
)

type FeedbackController interface {
}

type feedbackController struct {
	feedbackRepo repository.FeedbackRepository
}

func NewFeedbackRepository(feedbackRepo repository.FeedbackRepository) FeedbackController {
	return &feedbackController{feedbackRepo: feedbackRepo}
}

func (fc *feedbackController) CreateFeedback(ctx *fiber.Ctx) error {
	return nil
}

func (fc *feedbackController) GetFeedback(ctx *fiber.Ctx) error {
	return nil
}

func (fc *feedbackController) UpdateFeedback(ctx *fiber.Ctx) error {
	return nil
}

func (fc *feedbackController) DeleteFeedback(ctx *fiber.Ctx) error {
	return nil
}
