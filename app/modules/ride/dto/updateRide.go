package rideDto

import "time"

type UpdateRide struct {
	ID            uint   `json:"id"`
	Status        string `json:"status"`
	StartTime     time.Time
	EndTime       time.Time
	StartLocation string `gorm:"type:varchar(100)"`
	EndLocation   string `gorm:"type:varchar(100)"`
}

type UpdateRideResponse struct {
	ID            uint         `json:"id"`
	Status        string       `json:"status"`
	TransactionID uint         `json:"transaction_id"`
	Message       string       `json:"message"`
	Transaction   *Transaction `json:"transaction" `
	Rider         *Rider       `json:"rider"`
}

type Rider struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cards     []Card `json:"cards"`
}

type Card struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type Transaction struct {
	ID            uint   `json:"id"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Fare          uint   `json:"fare"`
}
