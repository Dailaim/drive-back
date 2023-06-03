package ride

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	RiderID       uint 
	DriverID      uint  
	StartLocation string `gorm:"type:varchar(100)"`
	EndLocation   string `gorm:"type:varchar(100)"`
	StartTime     time.Time
	EndTime       time.Time
	Status        string
	Fare          uint64
}

// "<Referencia><Monto><Moneda><SecretoIntegridad>"
//  "1-2490000-cop-publicKey"

// fmt.Sprintf("%d%d%s%s", ride.ID, ride.Fare, "cop", wompiKey.Integrity)
// encriptar con sha256
// data := 