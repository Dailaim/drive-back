package driver

import (
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	LastName         string
	FirstName        string
	Email           string `gorm:"type:varchar(100);unique_index"`
	CurrentLocation string       `gorm:"type:varchar(100)"`
	Rides           []ride.Model `gorm:"foreignKey:DriverID"`
}
