package seeds

import (
/* 	"github.com/Daizaikun/drive-back/app/lib" */
	cardSeed "github.com/Daizaikun/drive-back/app/modules/card/seed"
	driverSeed "github.com/Daizaikun/drive-back/app/modules/driver/seed"
	riderSeed "github.com/Daizaikun/drive-back/app/modules/rider/seed"

	"gorm.io/gorm"
)

func Run(DB *gorm.DB) {

	DB.Create(driverSeed.Seed)

	DB.Create(riderSeed.Seed)
	
	cardSeed.Seed[0].RiderID = riderSeed.Seed[0].ID
	cardSeed.Seed[1].RiderID = riderSeed.Seed[1].ID

	DB.Create(cardSeed.Seed)

}
