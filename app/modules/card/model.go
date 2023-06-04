package card

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name  string
	Token string
}
