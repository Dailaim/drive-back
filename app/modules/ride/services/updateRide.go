package rideServices

import (
	"strconv"
	"time"

	"github.com/Daizaikun/drive-back/app/db"
	"github.com/Daizaikun/drive-back/app/lib"
	rideDto "github.com/Daizaikun/drive-back/app/modules/ride/dto"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	riderModel "github.com/Daizaikun/drive-back/app/modules/rider/model"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary Update Ride
// @Description Update Ride with ID
// @Tags Ride
// @Accept json
// @Produce json
// @Param object body rideDto.UpdateRide true "Update Ride"
// @Success 200 {object} rideDto.UpdateRideResponse
// @Failure 400 {object} rideDto.ErrorResponse
// @Router /ride/{id} [put]
func UpdateRide(c *fiber.Ctx) error {

	id := c.Params("id")

	RideDto := rideDto.UpdateRide{}

	err := c.BodyParser(&RideDto)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	switch RideDto.Status {
	case "create":
		return updateCreateResponse(c, RideDto, id)
	case "processing":
		return updateProcessingResponse(c, RideDto, id)
	case "started":
		return updateStartedResponse(c, RideDto, id)
	case "completed":
		return updateCompletedResponse(c, RideDto, id)
	case "cancelled":
		return updateCancelledResponse(c, RideDto, id)
	default:
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}
}

func updateCompletedResponse(c *fiber.Ctx, RideDto rideDto.UpdateRide, id string) error {

	idUnit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var RideDB rideModel.Model

	RideDB.ID = uint(idUnit)
	RideDB.Status = "completed"
	RideDB.EndTime = time.Now()
	RideDB.EndLocation = RideDto.EndLocation

	result := db.Ctx.Model(&RideDB).Where("id = ?", RideDB.ID).Updates(&RideDB)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	RideStack, err := lib.CalculateRide(RideDB.StartLocation, RideDB.EndLocation, RideDB.StartTime, RideDB.EndTime)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: err.Error(),
			},
		},
		)
	}

	var transaction transactionModel.Model

	transaction.ID = RideDB.TransactionID
	transaction.Status = "pending"
	transaction.StatusMessage = "Transaction pending"
	transaction.Fare = uint(RideStack.Price)

	result = db.Ctx.Model(&transaction).Where("id = ?", transaction.ID).Updates(&transaction)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var Rider riderModel.Model

	result = db.Ctx.Preload("cards").First(&Rider, RideDB.RiderID)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	return c.JSON(rideDto.UpdateRideResponse{
		ID:            RideDB.ID,
		Status:        RideDB.Status,
		TransactionID: RideDB.TransactionID,
		Message:       "Ride completed successfully",
		Transaction: &rideDto.Transaction{
			ID:            RideDB.Transaction.ID,
			Status:        transaction.Status,
			Fare:          transaction.Fare,
			StatusMessage: transaction.StatusMessage,
		},
	})
}

func updateCancelledResponse(c *fiber.Ctx, RideDto rideDto.UpdateRide, id string) error {

	idUnit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var RideDB rideModel.Model

	RideDB.ID = uint(idUnit)
	RideDB.Status = "cancelled"

	db.Ctx.Model(&RideDB).Where("id = ?", RideDB.ID).Updates(&RideDB)

	return c.JSON(rideDto.UpdateRideResponse{
		Message: "Ride cancelled",
	})
}

func updateStartedResponse(c *fiber.Ctx, RideDto rideDto.UpdateRide, id string) error {
	idUnit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	var RideDB rideModel.Model

	RideDB.ID = uint(idUnit)
	RideDB.Status = "started"
	RideDB.StartTime = RideDto.StartTime
	RideDB.StartLocation = RideDto.StartLocation

	result := db.Ctx.Model(&RideDB).Where("id = ?", RideDto.ID).Updates(&RideDB)
	if result.Error != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid request",
			},
		},
		)
	}

	return c.JSON(rideDto.UpdateRideResponse{
		ID:      RideDB.ID,
		Status:  RideDB.Status,
		Message: "Ride created successfully",
	})
}

func updateProcessingResponse(c *fiber.Ctx, RideDto rideDto.UpdateRide, id string) error {

	idUni64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid ID request",
			},
		},
		)
	}

	var RideDB rideModel.Model
	RideDB.ID = uint(idUni64)
	RideDB.Status = "processing"

	return c.JSON(rideDto.UpdateRideResponse{
		ID:      RideDB.ID,
		Status:  RideDB.Status,
		Message: "Ride processing successfully",
	})
}

func updateCreateResponse(c *fiber.Ctx, RideDto rideDto.UpdateRide, id string) error {

	idUni64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(400).JSON(rideDto.ErrorResponse{
			Error: rideDto.Message{
				Message: "Invalid ID request",
			},
		},
		)
	}

	var RideDB rideModel.Model

	RideDB.ID = uint(idUni64)
	RideDB.Status = "processing"

	return c.JSON(rideDto.UpdateRideResponse{
		ID:      RideDB.ID,
		Status:  RideDB.Status,
		Message: "Ride created successfully",
	})
}
