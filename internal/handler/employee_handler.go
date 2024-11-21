package handler

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"
	"encoding/json"

	"github.com/alexliesenfeld/health"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type EmployeeHandler struct {
	service adapters.EmployeeService
}

func NewEmployeeHandler(service adapters.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (c *EmployeeHandler) CreateEmployee(context *fiber.Ctx) error {
	var employee entity.Employee
	if err := context.BodyParser(&employee); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": "invalid request",
		})
	}

	if err := c.service.CreateEmployee(&employee); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Create Employee Success",
	})
}

func (c *EmployeeHandler) GetEmployee(context *fiber.Ctx) error {
	employee, err := c.service.GetEmployee()
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(employee)
}

func (c *EmployeeHandler) GetEmployeeById(context *fiber.Ctx) error {
	id := context.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}

	employee, err := c.service.GetEmployeeById(uid)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}
	return context.Status(fiber.StatusOK).JSON(employee)
}

func (c *EmployeeHandler) UpdateEmployee(context *fiber.Ctx) error {
	id := context.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}

	var employee entity.Employee
	if err := context.BodyParser(&employee); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": "invalid request",
		})
	}

	if err := c.service.UpdateEmployee(uid, &employee); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   false,
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Update Employee Success",
	})
}

// NewHealthHandler is a Fiber handler for returning health status as JSON.
func NewHealthHandler(context *fiber.Ctx, checker health.Checker) error {
	// Run the health check and retrieve the status
	status := checker.Check(context.Context())

	// Serialize the status to JSON
	response, err := json.Marshal(status)
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).SendString("Failed to serialize health check status")
	}

	// Set JSON content-type and send the response
	context.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return context.Send(response)
}
