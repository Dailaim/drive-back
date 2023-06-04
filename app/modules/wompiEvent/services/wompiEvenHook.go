package services

import (
	"github.com/Daizaikun/drive-back/app/modules/wompiEvent/dto"
	"github.com/gofiber/fiber/v2"
)

func WompiEvenHook(c *fiber.Ctx) error {
	event := new(dto.Event)
	err := c.BodyParser(&event)

	if err != nil {
		return c.SendStatus(400)
	}


	return c.SendStatus(200)
}
