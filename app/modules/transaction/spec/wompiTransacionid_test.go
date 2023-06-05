package transaction_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	transactionDto "github.com/Daizaikun/drive-back/app/modules/transaction/dto"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	transactionServices "github.com/Daizaikun/drive-back/app/modules/transaction/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestWompiTransacionid(t *testing.T) {
	setup()

	// Create a transaction in db
	trans := transactionModel.Model{

		WompiTransactionID: "123",
		Status:             "completed",
		StatusMessage:      "Transaction finished successfully",
		Fare:               500,
	}
	mockDB.Create(&trans)

	defer mockDB.Delete(&trans)

	direction := fmt.Sprintf("/transaction/id/%v", trans.WompiTransactionID)

	// Create request
	req, _ := http.NewRequest("GET", direction, nil)

	// Create fiber app
	app := fiber.New()

	app.Get("/transaction/id/:id", transactionServices.WompiTransactionByID)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// converter body
	var body transactionDto.TransactionResponse

	json.NewDecoder(resp.Body).Decode(&body)

	assert.Equal(t, trans.Status, body.Status)

	assert.Equal(t, trans.WompiTransactionID, body.WompiID)
}
