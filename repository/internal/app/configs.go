package app

import "github.com/gofiber/fiber/v3"

type Configs struct {
	Port     string       `json:"http_port"`
	BasePath string       `json:"base_path"`
	App      fiber.Config `json:"app"`
}
