package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	app2 "project/internal/app"
)

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
