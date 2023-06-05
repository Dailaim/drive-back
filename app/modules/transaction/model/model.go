package transactionModel

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	WompiTransactionID string
	Status             string
	StatusMessage      string
	Fare               uint
}

func (Model) TableName() string {
	return "transactions"
}
