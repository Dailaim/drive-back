package routers

import (
	"github.com/Daizaikun/drive-back/app/modules/wompiEvent"
	"github.com/gofiber/fiber/v2"
)

func List(app *fiber.App) {

	app.Mount("/", wompiEvent.Controller())
	

}
