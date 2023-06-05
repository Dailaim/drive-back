package db

import (
	"github.com/Daizaikun/drive-back/app/db/dns"
	"github.com/Daizaikun/drive-back/app/db/migrate"

	"github.com/Daizaikun/drive-back/app/db/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	/* 	"gorm.io/gorm/logger" */)

	
var Ctx *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.Open(dns.New()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	/* DB.Logger = logger.Default.LogMode(logger.Info) */

	err = migrate.New(db)
	if err != nil {
		panic("failed to migrate database")
	}

	seeds.Run(db)

	Ctx = db

}
