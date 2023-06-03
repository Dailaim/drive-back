package wompiEvent

import (
	"encoding/json"
	"fmt"

	"github.com/Daizaikun/drive-back/app/modules/wompiEvent/dto"
	"github.com/gofiber/fiber/v2"
)

// Path: app\modules\wompiEvent\controllers

func Controller() *fiber.App {
	app := fiber.New()

	app.Post("/wompievent", func(c *fiber.Ctx) error {
		event := new(dto.Event)
		err := c.BodyParser(&event)

		if err != nil {
			return c.SendStatus(400)
		}

		// converter to json
		json, err := json.Marshal(event)

		if err != nil {
			return c.SendStatus(400)
		}

		fmt.Println(string(json))

		return c.SendStatus(200)
	})
	return app
}
