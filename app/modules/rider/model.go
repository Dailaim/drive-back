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
	CardID    []uint
	Card      []card.Model `gorm:"references:CardID"`
	Rides     []ride.Model `gorm:"foreignKey:RiderID"`
}
