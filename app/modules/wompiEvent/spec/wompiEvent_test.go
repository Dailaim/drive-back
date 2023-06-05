package wompi_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Daizaikun/drive-back/app/db"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	wompiEventDto "github.com/Daizaikun/drive-back/app/modules/wompiEvent/dto"
	wompiServices "github.com/Daizaikun/drive-back/app/modules/wompiEvent/services"
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

func TestWompiEvenHook(t *testing.T) {
	setup()

	// Create test event
	testEvent := wompiEventDto.Event{
		Event: "testEvent",
		Data: wompiEventDto.Data{
			Transaction: wompiEventDto.Transaction{
				ID:        "123",
				Reference: "456",
				Status:    "success",
			},
		},
	}

	jsonEvent, _ := json.Marshal(testEvent)

	// Create request
	req, _ := http.NewRequest("POST", "/wompievent", bytes.NewBuffer(jsonEvent))
	req.Header.Set("Content-Type", "application/json")


	// Create fiber app
	app := fiber.New()

	app.Post("/wompievent", wompiServices.WompiEvenHook)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check database
	var transaction transactionModel.Model
	mockDB.Where("wompi_transaction_id = ?", testEvent.Data.Transaction.ID).First(&transaction)

	assert.Equal(t, testEvent.Data.Transaction.ID, transaction.WompiTransactionID)
	assert.Equal(t, testEvent.Data.Transaction.Status, transaction.Status)
}
