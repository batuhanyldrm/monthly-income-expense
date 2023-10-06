package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

type Config struct {
	AppPort         int
	Host            string
	DBReplicaSetUrl string
}

func main() {
	fmt.Println("Monthly Income Expense service sarting...")

	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)

	fmt.Println("Monthly Income Expense service started at 8080 ...")
	app.Listen(":8080")
}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
	}))

	app.Get("/salaries", api.GetSalariesHandler)
	app.Get("/salary/:id", api.GetSalaryHandler)
	app.Post("/salary", api.PostSalaryHandler)
	app.Delete("/salary/:id", api.DeleteSalaryHandler)

	return app
}
