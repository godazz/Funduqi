package api

import (
	"github.com/godazz/Funduqi/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Ahmed",
		LastName:  "Abd El Gawad",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Ahmed")
}
