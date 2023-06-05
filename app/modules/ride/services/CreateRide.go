package rideServices

import (
	"github.com/Daizaikun/drive-back/app/db"
	rideDto "github.com/Daizaikun/drive-back/app/modules/ride/dto"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	"github.com/gofiber/fiber/v2"
)

func CreateRide(c *fiber.Ctx) error {

	DB := db.New()

	RideDto := new(rideDto.CreateRide)

	err := c.BodyParser(RideDto)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	RideDB := rideModel.Model{
		DriverID: RideDto.DriverID,
		RiderID:  RideDto.RiderID,
	}

	DB.Create(&RideDB)

	return c.JSON(rideDto.CreateRideResponse{
		ID: RideDB.ID,
	})
}
