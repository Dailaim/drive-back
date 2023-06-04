package migrate

import (
	"github.com/Daizaikun/drive-back/app/modules/card"
	"github.com/Daizaikun/drive-back/app/modules/driver"
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"github.com/Daizaikun/drive-back/app/modules/rider"
	"github.com/Daizaikun/drive-back/app/modules/transaction"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) error {

	err := DB.AutoMigrate(
		&transaction.Model{},
		&driver.Model{},
		&rider.Model{},
		&ride.Model{},
		&card.Model{},
	)

	return err
}
