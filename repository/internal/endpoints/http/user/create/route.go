package create

import "github.com/gofiber/fiber/v3"

// Endpoint short-description
// @Summary     summary
// @Description Long-description
// @Tags        Users
// @Produce     json
// @Accept 		json
// @Param       body body Params  true "Request Params"
// @Success     200  {object} Response
// @Failure     default  {object} fiber.Error
// @Router      /api/users [post]
func Endpoint() *fiber.Route {
	return &fiber.Route{
		Method:   "POST",
		Handlers: Handlers,
		Path:     "/users",
	}
}
