package transaction

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	WompiTransactionID string
	Status             string
	StatusMessage      string
	Fare               uint
}