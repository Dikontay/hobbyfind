package create

import "github.com/gofiber/fiber/v3"

// Endpoint short-description
// @Summary     summary
// @Description Long-description
// @Tags        Users
// @Produce     json
// @Accept 		json
// @Param       body body params.Main  true "Request Params"
// @Success     200  {object} info.Main
// @Failure     default  {object} responses.ErrorResponse
// @Router      /api/create/ [post]

func Route() *fiber.Route {
	return &fiber.Route{
		Method:   "POST",
		Handlers: Handlers,
		Path:     "/api/create/",
	}
}
