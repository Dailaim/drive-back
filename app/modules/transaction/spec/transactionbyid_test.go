package transaction_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Daizaikun/drive-back/app/db"
	transactionDto "github.com/Daizaikun/drive-back/app/modules/transaction/dto"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	transactionServices "github.com/Daizaikun/drive-back/app/modules/transaction/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var mockDB *gorm.DB

func setup() {
	DB := db.Connect()

	mockDB = DB

	mockDB.AutoMigrate(&transactionModel.Model{})
}

func TestTransactionByID(t *testing.T) {
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

	direction := fmt.Sprintf("/transaction/id/%d", trans.ID)

	// Create request
	req, _ := http.NewRequest("GET", direction, nil)

	// Create fiber app
	app := fiber.New()

	app.Get("/transaction/id/:id", transactionServices.TransactionByID)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// converter body
	var body transactionDto.TransactionResponse

	json.NewDecoder(resp.Body).Decode(&body)

	assert.Equal(t, trans.Status, body.Status)

	assert.Equal(t, trans.WompiTransactionID, body.WompiID)
}
