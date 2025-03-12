package main

import (
	"fmt"
	app2 "github.com/Dikontay/hobbyfind/repository/internal/app"
	"github.com/gofiber/fiber/v3"
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
