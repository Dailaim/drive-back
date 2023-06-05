package rideServices

import (
	"github.com/Daizaikun/drive-back/app/db"
	rideDto "github.com/Daizaikun/drive-back/app/modules/ride/dto"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create Ride
// @Description Create Ride with DriverID and RiderID
// @Tags Ride
// @Accept json
// @Produce json
// @Param object body rideDto.CreateRide true "Create Ride"
// @Success 200 {object} rideDto.CreateRideResponse
// @Failure 400 {object} rideDto.ErrorResponse
// @Router /ride [post]
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
	RideDB.Status = "created"

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
		ID:     RideDB.ID,
		Status: RideDB.Status,
	})
}
