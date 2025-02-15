package main

import (
	"FeedbackAPI/config"
	"FeedbackAPI/controllers"
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

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello"})
	})

	app.Post("/signin", customerController.SignInCustomer)
	app.Post("/signup", customerController.SignUpCustomer)

	_ = app.Listen(":8080")

}
