package rideModel

import (
	"time"

	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	DriverID      uint
	RiderID       uint
	StartLocation string `gorm:"type:varchar(100)"`
	EndLocation   string `gorm:"type:varchar(100)"`
	StartTime     time.Time
	EndTime       time.Time
	Status        string
	TransactionID string
	Transaction   transactionModel.Model
}

func (Model) TableName() string {
	return "rides"
}
