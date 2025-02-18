package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger"
)

type Configs struct {
	Port     string         `json:"http_port"`
	BasePath string         `json:"base_path"`
	App      fiber.Config   `json:"app"`
	Swagger  SwaggerConfigs `json:"swagger"`
}

type SwaggerConfigs struct {
	Path    string         `json:"path"`
	Configs swagger.Config `json:"configs" config:"ignore"`
}
