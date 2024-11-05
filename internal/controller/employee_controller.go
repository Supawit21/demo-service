package controller

import (
	"Supawit21/demo_service/internal/adapters"
	"Supawit21/demo_service/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	service adapters.EmployeeService
}

func NewEmployeeController(service adapters.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) CreateEmployee(context *fiber.Ctx) error {
	var employee entity.Employee
	if err := context.BodyParser(&employee); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := c.service.CreateEmployee(&employee); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return context.Status(fiber.StatusCreated).JSON(employee)

}
