package main

import (
	"Supawit21/demo_service/internal/controller"
	"Supawit21/demo_service/internal/repository"
	"Supawit21/demo_service/internal/service"
	"Supawit21/demo_service/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	employeedb := database.InitialDatabase()
	employeeRepo := repository.NewEmployeeRepository(employeedb)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	app.Post("/employee", employeeController.CreateEmployee)

	log.Fatal(app.Listen(":8888"))
}
