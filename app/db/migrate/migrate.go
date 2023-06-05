package migrate

import (
	cardModel "github.com/Daizaikun/drive-back/app/modules/card/model"
	driverModel "github.com/Daizaikun/drive-back/app/modules/driver/model"
	rideModel "github.com/Daizaikun/drive-back/app/modules/ride/model"
	riderModel "github.com/Daizaikun/drive-back/app/modules/rider/model"
	transactionModel "github.com/Daizaikun/drive-back/app/modules/transaction/model"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) error {

	err := DB.AutoMigrate(
		&transactionModel.Model{},
		&driverModel.Model{},
		&riderModel.Model{},
		&rideModel.Model{},
		&cardModel.Model{},
	)

	return err
}
