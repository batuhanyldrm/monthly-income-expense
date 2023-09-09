package main

import "github.com/gofiber/fiber"

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
