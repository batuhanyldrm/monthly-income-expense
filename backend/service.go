package main

import (
	"monthly-income-expense/models"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) Service {
	return Service{
		repository: repository,
	}
}

func (service *Service) GetSalaries() ([]models.Salary, error) {

	salaries, err := service.repository.GetSalaries()

	if err != nil {
		return nil, err
	}

	return salaries, nil
}

func (service *Service) GetSalary(ID string) (models.Salary, error) {
	salary, err := service.repository.GetSalary(ID)

	if err != nil {
		return models.Salary{}, nil
	}

	return salary, nil
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}

	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
