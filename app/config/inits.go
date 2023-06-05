package config

import (
	"github.com/Daizaikun/drive-back/app/db"
)

func Inits() {

	db.Connect()

}
