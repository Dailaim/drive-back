package migrate

import (

	"gorm.io/gorm"
)

func New(DB *gorm.DB) error {

	err := DB.AutoMigrate(

	)

	return err
}
