package ride

import (
	rideServices "github.com/Daizaikun/drive-back/app/modules/ride/services"
	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	
	app := fiber.New()

	app.Post("/ride", rideServices.CreateRide)

	app.Put("/ride/:id", rideServices.UpdateRide)

	return app
}
