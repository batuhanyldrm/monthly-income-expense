package main

import (
	"monthly-income-expense/models"

	"github.com/gofiber/fiber"
)

type Api struct {
	service *Service
}

func NewApi(service *Service) Api {
	return Api{
		service: service,
	}
}

func (api *Api) GetSalariesHandler(c *fiber.Ctx) {

	salaries, err := api.service.GetSalaries()

	switch err {
	case nil:
		c.JSON(salaries)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (api *Api) GetSalaryHandler(c *fiber.Ctx) {
	ID := c.Params("id")

	salary, err := api.service.GetSalary(ID)

	switch err {
	case nil:
		c.JSON(salary)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (api *Api) PostSalaryHandler(c *fiber.Ctx) {
	createSalary := models.SalaryDTO{}
	err := c.BodyParser(&createSalary)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	salary := api.service.PostSalary(createSalary)

	switch err {
	case nil:
		c.JSON(salary)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}
