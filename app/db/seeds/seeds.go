package seeds

import (
	"github.com/Daizaikun/drive-back/app/lib"
	"github.com/Daizaikun/drive-back/app/modules/card"
	"github.com/Daizaikun/drive-back/app/modules/driver"
	"github.com/Daizaikun/drive-back/app/modules/rider"
	"gorm.io/gorm"
)

func Run(DB *gorm.DB) {

	if lib.IsProduction() {
		return
	}

	DB.Create(driver.Seed)

	DB.Create(rider.Seed)
	card.Seed[0].RiderID = rider.Seed[0].ID
	card.Seed[0].RiderID = rider.Seed[1].ID


	DB.Create(card.Seed)

}
