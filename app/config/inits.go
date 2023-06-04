package config

import (
	"github.com/Daizaikun/drive-back/app/db"
	"github.com/Daizaikun/drive-back/app/lib"
)

func Inits() {

	db.Connect()

	(&lib.WompiKey{}).SetWompiKey()
	
}
