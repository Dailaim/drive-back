package rider

import (
	"github.com/Daizaikun/drive-back/app/modules/card"
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	LastName  string
	FirstName string
	Card      []card.Model `gorm:"foreignKey:RiderID"`
	Rides     []ride.Model `gorm:"foreignKey:RiderID"`
}

func (Model) TableName() string {
	return "riders"
}


