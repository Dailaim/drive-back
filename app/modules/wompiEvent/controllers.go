package wompiEvent

import (
	"github.com/Daizaikun/drive-back/app/modules/wompiEvent/services"
	"github.com/gofiber/fiber/v2"
)

// Path: app\modules\wompiEvent\controllers

func Controller() *fiber.App { 
	app := fiber.New()

	app.Post("/wompievent", wompiServices.WompiEvenHook)
	return app
}
