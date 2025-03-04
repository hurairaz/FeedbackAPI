package main

import (
	"FeedbackAPI/config"
	"FeedbackAPI/controllers"
	"FeedbackAPI/middleware"
	"FeedbackAPI/repository"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//if err := config.MigrateModels(db); err != nil {
	//	log.Fatalf("Migration Failed: %v", err)
	//}

	app := fiber.New()

	customerRepo := repository.NewCustomerRepository(db)
	customerController := controllers.NewCustomerController(customerRepo)
	feedbackRepo := repository.NewFeedbackRepository(db)
	feedbackController := controllers.NewFeedbackController(feedbackRepo)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello"})
	})

	app.Post("/signin", customerController.SignInCustomer)
	app.Post("/signup", customerController.SignUpCustomer)
	app.Get("/customers/:id", middleware.AuthMiddleware, customerController.GetCustomer)
	app.Put("/customers/:id", middleware.AuthMiddleware, customerController.UpdateCustomer)
	app.Get("/customers/:id/feedbacks", middleware.AuthMiddleware, customerController.GetCustomerFeedbacks)
	app.Delete("/customers/:id", middleware.AuthMiddleware, customerController.DeleteCustomer)
	app.Post("/feedbacks/", feedbackController.CreateFeedback)
	app.Get("/feedbacks/:id", feedbackController.GetFeedback)
	app.Put("/feedbacks/:id", feedbackController.UpdateFeedback)
	app.Delete("/feedbacks/:id", feedbackController.DeleteFeedback)
	_ = app.Listen(":8080")

}
