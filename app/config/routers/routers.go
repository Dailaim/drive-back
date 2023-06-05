package routers

import (
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"github.com/Daizaikun/drive-back/app/modules/wompiEvent"
	"github.com/gofiber/fiber/v2"
)

func List(app *fiber.App) {

	app.Mount("/", wompiEvent.Controller())

	app.Mount("ride" , ride.Controller())
	
}
