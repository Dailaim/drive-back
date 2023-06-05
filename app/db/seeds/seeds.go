package seeds

import (
	"github.com/Daizaikun/drive-back/app/lib"
	cardSeed "github.com/Daizaikun/drive-back/app/modules/card/seed"
	driverSeed "github.com/Daizaikun/drive-back/app/modules/driver/seed"
	riderSeed "github.com/Daizaikun/drive-back/app/modules/rider/seed"

	"gorm.io/gorm"
)

func Run(DB *gorm.DB) {

	if lib.IsProduction() {
		return
	}

	result := DB.Create(driverSeed.Seed)
	if result.Error != nil {
		panic("Error creating cards")
	}

	result = DB.Create(riderSeed.Seed)
	if result.Error != nil {
		panic("Error creating cards")
	}

	RidersID := []uint{riderSeed.Seed[0].ID, riderSeed.Seed[1].ID}

	cards := cardSeed.GenerateSeed(RidersID)

	result = DB.Create(&cards)

	if result.Error != nil {
		panic("Error creating cards")
	}

}
