package main

import (
	"Supawit21/demo_service/internal/handler"
	"Supawit21/demo_service/internal/repository"
	"Supawit21/demo_service/internal/service"
	"Supawit21/demo_service/pkg/database"
	"Supawit21/demo_service/pkg/utils"
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	employeeDB := database.InitialDatabase()
	employeeRepo := repository.NewEmployeeRepository(employeeDB)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// healthcheck
	healthCheck := utils.HealthCheck(employeeDB)

	app := fiber.New()
	api := app.Group("/api")

	apiVersion := api.Group("/v1")
	apiVersion.Get("/employee", employeeHandler.GetEmployee)
	apiVersion.Get("/employee/:id", employeeHandler.GetEmployeeById)
	apiVersion.Post("/employee", employeeHandler.CreateEmployee)
	apiVersion.Put("/employee/:id", employeeHandler.UpdateEmployee)

	// endpoint healthcheck
	api.Get("/health", func(c *fiber.Ctx) error {
		return handler.NewHealthHandler(c, healthCheck)
	})

	ln, _ := net.Listen("tcp", ":"+os.Getenv("APP_PORT"))
	app.Listener(ln)
}
