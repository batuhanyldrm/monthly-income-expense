package main

import (
	"bytes"
	"encoding/json"
	"io"
	"monthly-income-expense/models"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSalaries(t *testing.T) {
	Convey("Get Salaries", t, func() {
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
			req, _ := http.NewRequest(http.MethodGet, "/salaries", nil)
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

				//So(actualResult, ShouldHaveLength, 2)
				//So(actualResult[0].ID, ShouldEqual, salary.ID)
				So(actualResult[0].Debit, ShouldEqual, salary.Debit)
				So(actualResult[0].Salary, ShouldEqual, salary.Salary)
				So(actualResult[0].MoneyGain, ShouldEqual, salary.MoneyGain)
				//So(actualResult[0].CreatedAt, ShouldEqual, salary.CreatedAt)
				//So(actualResult[0].UpdatedAt, ShouldEqual, salary.UpdatedAt)
				//So(actualResult[0].Users[0].ID, ShouldEqual, salary.Users[0].ID)
				So(actualResult[0].Users[0].Name, ShouldEqual, salary.Users[0].Name)
				So(actualResult[0].Users[0].Email, ShouldEqual, salary.Users[0].Email)
				//So(actualResult[0].Users[0].UpdatedAt, ShouldEqual, salary.Users[0].UpdatedAt)
				So(actualResult[1].Debit, ShouldEqual, salary1.Debit)
				So(actualResult[1].Salary, ShouldEqual, salary1.Salary)
				So(actualResult[1].MoneyGain, ShouldEqual, salary1.MoneyGain)
				//So(actualResult[1].CreatedAt, ShouldEqual, salary.CreatedAt)
				//So(actualResult[1].UpdatedAt, ShouldEqual, salary.UpdatedAt)
				//So(actualResult[1].Users[0].ID, ShouldEqual, salary.Users[0].ID)
				So(actualResult[1].Users[0].Name, ShouldEqual, salary1.Users[0].Name)
				So(actualResult[1].Users[0].Email, ShouldEqual, salary1.Users[0].Email)
				//So(actualResult[1].Users[0].UpdatedAt, ShouldEqual, salary.Users[0].UpdatedAt)
			})
		})
	})
}

func TestGetSalary(t *testing.T) {
	Convey("Get Salary", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		salary := models.Salary{
			ID:        GenerateUUID(8),
			Salary:    "2000",
			Debit:     "100",
			MoneyGain: "200",
			CreatedAt: time.Now().UTC().Round(time.Second),
			UpdatedAt: time.Now().UTC().Round(time.Second),
			Users: []models.User{
				models.User{
					ID:        GenerateUUID(8),
					Name:      "Metehan",
					Email:     "metehan@gmail.com",
					CreatedAt: time.Now().UTC().Round(time.Second),
					UpdatedAt: time.Now().UTC().Round(time.Second),
				}},
		}

		repository.CreateSalary(salary)

		Convey("When the get request sent", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/salary/"+salary.ID, nil)
			resp, err := app.Test(req, 3000)

			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product should be returned", func() {
				actualResult := models.Salary{}
				actualRespBody, _ := io.ReadAll(resp.Body)
				err := json.Unmarshal(actualRespBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult.ID, ShouldEqual, salary.ID)
				So(actualResult.Salary, ShouldEqual, salary.Salary)
				So(actualResult.Debit, ShouldEqual, salary.Debit)
				So(actualResult.MoneyGain, ShouldEqual, salary.MoneyGain)
				So(actualResult.UpdatedAt, ShouldEqual, salary.UpdatedAt)
				So(actualResult.CreatedAt, ShouldEqual, salary.CreatedAt)
				So(actualResult.Users[0].ID, ShouldEqual, salary.Users[0].ID)
				So(actualResult.Users[0].Name, ShouldEqual, salary.Users[0].Name)
				So(actualResult.Users[0].Email, ShouldEqual, salary.Users[0].Email)
				So(actualResult.Users[0].CreatedAt, ShouldEqual, salary.Users[0].CreatedAt)
				So(actualResult.Users[0].UpdatedAt, ShouldEqual, salary.Users[0].UpdatedAt)

			})
		})
	})
}

func TestAddSalary(t *testing.T) {
	Convey("Add stock", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Salary{
			Salary:    "1222",
			Debit:     "122",
			MoneyGain: "23",
			Users: []models.User{models.User{
				Name:      "batu",
				Email:     "aa@gmail.com",
				CreatedAt: time.Now().UTC().Round(time.Second),
				UpdatedAt: time.Now().UTC().Round(time.Second),
			}},
			CreatedAt: time.Now().UTC().Round(time.Second),
			UpdatedAt: time.Now().UTC().Round(time.Second),
		}

		Convey("When the post request sent", func() {

			reqBody, err := json.Marshal(stock)

			req, _ := http.NewRequest(http.MethodPost, "/salary", bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

			app := SetupApp(&api)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 201", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusCreated)
			})

			Convey("Then added stock should return", func() {
				actualResult, err := repository.GetSalary(stock.ID)

				So(err, ShouldBeNil)
				So(actualResult, ShouldNotBeNil)
				So(actualResult.Salary, ShouldEqual, stock.Salary)
				So(actualResult.Debit, ShouldEqual, stock.Debit)
				So(actualResult.MoneyGain, ShouldEqual, stock.MoneyGain)
			})
		})
	})
}
