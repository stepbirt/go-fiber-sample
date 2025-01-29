package handler

import (
	"gofiber/service"

	"github.com/gofiber/fiber/v2"
)

type userhandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) userhandler {
	return userhandler{userService: service}
}

func (h *userhandler) SignUp(c *fiber.Ctx) error {

	request := service.RequestNewUser{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	result, err := h.userService.NewUser(request)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "unexcpeted error",
		})
		return err
	}
	c.Status(fiber.StatusOK)
	return c.JSON(result)
}
