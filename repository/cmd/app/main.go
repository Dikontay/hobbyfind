package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	_ "github.com/gofiber/swagger"
	_ "project/cmd/docs"
	app2 "project/internal/app"
)

// @title HobbyLink API
// @version 1.0
// @description User Repository.
// @termsOfService http://swagger.io/terms/
// @BasePath /api/
func main() {
	app := app2.NewService(app2.Configs{
		Port:     ":8080",
		BasePath: "/api",
		App:      fiber.Config{},
	})

	err := app.Start()
	if err != nil {
		fmt.Println(err)
	}
}
