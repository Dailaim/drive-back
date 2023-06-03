package rider

import (
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Name          string
	Email         string `gorm:"type:varchar(100);unique_index"`
	Password      string
	PaymentSource string
	Rides         []ride.Model `gorm:"foreignKey:RiderID"`
}
