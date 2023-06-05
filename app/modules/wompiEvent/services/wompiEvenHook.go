package wompiServices

import (
	"strconv"

	"github.com/Daizaikun/drive-back/app/db"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	wompiEventDto "github.com/Daizaikun/drive-back/app/modules/wompiEvent/dto"
	"github.com/gofiber/fiber/v2"
)


// @Summary Wompi Event Hook
// @Description Wompi Event Hook
// @Tags WompiEvent
// @Accept json
// @Produce json
// @Param object body wompiEventDto.Event true "Wompi Event"
// @Success 200
// @Failure 400 {object} wompiEventDto.ErrorResponse
// @Router /wompievent [post]
func WompiEvenHook(c *fiber.Ctx) error {

	event := new(wompiEventDto.Event)

	err := c.BodyParser(&event)

	if err != nil {
		return c.Status(400).JSON(&wompiEventDto.ErrorResponse{
			Error: wompiEventDto.Message{
				Message: "Invalid request body",
			},
		})
	}

	var transaction transactionModel.Model

	transactionID, err := strconv.ParseUint(event.Data.Transaction.Reference, 10, 32)

	if err != nil {
		return c.Status(400).JSON(&wompiEventDto.ErrorResponse{
			Error: wompiEventDto.Message{
				Message: "Invalid transaction reference",
			},
		})

	}

	transaction.ID = uint(transactionID)

	transaction.WompiTransactionID = event.Data.Transaction.ID

	transaction.Status = event.Data.Transaction.Status

	transaction.StatusMessage = "Transaction finished successfully"

	result := db.Ctx.Save(&transaction)
	if result.Error != nil {
		return c.Status(500).JSON(&wompiEventDto.ErrorResponse{
			Error: wompiEventDto.Message{
				Message: "Error saving transaction",
			},
		})

	}

	return c.SendStatus(200)
}
