package transactionServices

import (
	"github.com/Daizaikun/drive-back/app/db"
	transactionDto "github.com/Daizaikun/drive-back/app/modules/transaction/dto"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get Wompi Transaction by ID
// @Description Get Wompi Transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Wompi Transaction ID"
// @Success 200 {object} transactionDto.TransactionResponse
// @Failure 400 {object} transactionDto.ErrorResponse
// @Router /transaction/wompi/{id} [get]
func WompiTransactionByID(c *fiber.Ctx) error {

	WompiID := c.Params("id")

	transaction := new(transactionModel.Model)

	result := db.Ctx.Model(&transactionModel.Model{}).Where("wompi_transaction_id = ?", WompiID).First(&transaction)
	if result.Error != nil {
		return c.Status(400).JSON(&transactionDto.ErrorResponse{
			Error: transactionDto.Message{
				Message: "Invalid wompi transaction id",
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
