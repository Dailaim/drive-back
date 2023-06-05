package transactionServices

import (
	"github.com/Daizaikun/drive-back/app/db"
	transactionDto "github.com/Daizaikun/drive-back/app/modules/transaction/dto"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"github.com/gofiber/fiber/v2"
)


// @Summary Get Transaction by ID
// @Description Get Transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} transactionDto.TransactionResponse
// @Failure 400 {object} transactionDto.ErrorResponse
// @Router /transaction/id/{id} [get]
func TransactionByID(c *fiber.Ctx) error {

	transactionID := c.Params("id")

	transaction := new(transactionModel.Model)

	result := db.Ctx.First(&transaction, transactionID)
	if result.Error != nil {
		return c.Status(400).JSON(&transactionDto.ErrorResponse{
			Error: transactionDto.Message{
				Message: "Invalid transaction id",
			},
		})
	}

	return c.JSON(&transactionDto.TransactionResponse{
		ID:      transaction.ID,
		Status:  transaction.Status,
		Message: transaction.StatusMessage,
		Fare:    transaction.Fare,
		WompiID: transaction.WompiTransactionID,
	})
}