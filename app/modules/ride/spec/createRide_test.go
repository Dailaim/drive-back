package rideServices_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Daizaikun/drive-back/app/db"
	rideDto "github.com/Daizaikun/drive-back/app/modules/ride/dto"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	rideServices "github.com/Daizaikun/drive-back/app/modules/ride/services"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var mockDB *gorm.DB

func setup() {
	DB := db.Connect()

	mockDB = DB

	mockDB.AutoMigrate(&rideModel.Model{})
}

func TestCreateRide(t *testing.T) {
	setup()

	rideRequest := rideDto.CreateRide{
		DriverID: 1,
		RiderID:  1,
	}

	requestBody, _ := json.Marshal(rideRequest)

	// Create request
	req, _ := http.NewRequest("POST", "/ride", bytes.NewBuffer(requestBody))

	// Create fiber app
	app := fiber.New()

	app.Post("/ride", rideServices.CreateRide)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response
	var body rideDto.CreateRideResponse

	json.NewDecoder(resp.Body).Decode(&body)

	assert.Equal(t, "created", body.Status)
}
