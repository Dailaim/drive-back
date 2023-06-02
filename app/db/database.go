package db

import (
	"github.com/Daizaikun/wompi-back/app/db/dns"
	"github.com/Daizaikun/wompi-back/app/db/migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	/* 	"gorm.io/gorm/logger" */)

var DB *gorm.DB

func Connect() {

	DB, err := gorm.Open(postgres.Open(dns.New()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	/* DB.Logger = logger.Default.LogMode(logger.Info) */

	err = migrate.New(DB)
	if err != nil {
		panic("failed to migrate database")
	}

}
