package ride

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	RiderID       uint 
	DriverID      uint  
	StartLocation string `gorm:"type:varchar(100)"`
	EndLocation   string `gorm:"type:varchar(100)"`
	StartTime     time.Time
	EndTime       time.Time
	Status        string
	Fare          float64
}
