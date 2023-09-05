package main

import (
	"encoding/json"
	"io"
	"monthly-income-expense/models"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSalary(t *testing.T) {
	Convey("Get Salary", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		salary := models.Salary{
			ID:        GenerateUUID(8),
			Salary:    "17500",
			Debit:     "8000",
			MoneyGain: "500",
			CreatedAt: time.Now().UTC().Round(time.Second),
			UpdatedAt: time.Now().UTC().Round(time.Second),
			Users: []models.User{
				models.User{
					ID:    GenerateUUID(8),
					Name:  "Batuhan Yildirim",
					Email: "batu@gmail.com",
				},
			},
		}

		salary1 := models.Salary{
			ID:        GenerateUUID(8),
			Salary:    "20500",
			Debit:     "10000",
			MoneyGain: "1000",
			CreatedAt: time.Now().UTC().Round(time.Second),
			UpdatedAt: time.Now().UTC().Round(time.Second),
			Users: []models.User{
				models.User{
					ID:    GenerateUUID(8),
					Name:  "Ali Kural",
					Email: "ali@gmail.com",
				},
			},
		}

		repository.CreateSalary(salary)
		repository.CreateSalary(salary1)

		Convey("When the get request sent", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			resp, err := app.Test(req, 3000)

			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product should be returned", func() {
				actualResult := []models.Salary{}
				actualRespBody, _ := io.ReadAll(resp.Body)
				err := json.Unmarshal(actualRespBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult, ShouldHaveLength, 2)
				So(actualResult[0].ID, ShouldEqual, salary.ID)
				So(actualResult[0].Debit, ShouldEqual, salary.Debit)
				So(actualResult[0].Salary, ShouldEqual, salary.Salary)
				So(actualResult[0].MoneyGain, ShouldEqual, salary.MoneyGain)
				So(actualResult[0].CreatedAt, ShouldEqual, salary.CreatedAt)
				So(actualResult[0].UpdatedAt, ShouldEqual, salary.UpdatedAt)
				So(actualResult[0].Users[0].ID, ShouldEqual, salary.Users[0].ID)
				So(actualResult[0].Users[0].Name, ShouldEqual, salary.Users[0].Name)
				So(actualResult[0].Users[0].Email, ShouldEqual, salary.Users[0].Email)
				So(actualResult[0].Users[0].UpdatedAt, ShouldEqual, salary.Users[0].UpdatedAt)
			})
		})
	})
}
