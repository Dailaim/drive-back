package rideServices

import (
	"github.com/Daizaikun/drive-back/app/db"
	rideDto "github.com/Daizaikun/drive-back/app/modules/ride/dto"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"github.com/gofiber/fiber/v2"
)

func CreateRide(c *fiber.Ctx) error {

	RideDto := rideDto.CreateRide{}

	err := c.BodyParser(&RideDto)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var transaction transactionModel.Model

	result := db.Ctx.Create(&transaction)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var RideDB rideModel.Model

	RideDB.DriverID = RideDto.DriverID
	RideDB.RiderID = RideDto.RiderID
	RideDB.TransactionID = transaction.ID

	result = db.Ctx.Create(&RideDB)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	return c.JSON(rideDto.CreateRideResponse{
		ID: RideDB.ID,
	})
}
