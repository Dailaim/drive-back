package config

import (
	"github.com/Daizaikun/drive-back/app/db"
	"github.com/Daizaikun/drive-back/app/lib"
)

func init() {
	db.Connect()
}

func Inits() {
	(&lib.WompiKey{}).SetWompiKey()
}
