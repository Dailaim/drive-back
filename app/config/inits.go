package config

import (
	"github.com/Daizaikun/wompi-back/app/db"
	"github.com/Daizaikun/wompi-back/app/lib"
)

func init() {
	db.Connect()
}

func Inits() {
	(&lib.WompiKey{}).SetWompiKey()

}
