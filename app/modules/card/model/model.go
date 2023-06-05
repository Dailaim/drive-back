package cardModel

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name   string
	Token  string
	RiderID uint 
}

func (Model) TableName() string {
	return "cards"
}
