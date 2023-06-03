package migrate

import (
	"github.com/Daizaikun/drive-back/app/modules/driver"
	"github.com/Daizaikun/drive-back/app/modules/ride"
	"github.com/Daizaikun/drive-back/app/modules/rider"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) error {

	err := DB.AutoMigrate(
		&driver.Model{},
		&ride.Model{},
		&rider.Model{},
	)

	return err
}
