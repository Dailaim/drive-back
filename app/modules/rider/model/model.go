package riderModel

import (
	cardModel "github.com/Daizaikun/drive-back/app/modules/card/model"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	LastName  string
	FirstName string
	Card      []cardModel.Model `gorm:"foreignKey:RiderID"`
	Rides     []rideModel.Model `gorm:"foreignKey:RiderID"`
}

func (Model) TableName() string {
	return "riders"
}


