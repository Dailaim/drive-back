package driverSeed

import "github.com/Daizaikun/drive-back/app/modules/driver/model"

var Seed = []*driverModel.Model{
	{
		LastName:        "Doe",
		FirstName:       "John",
		Email:           "driver1@example.com",
		CurrentLocation: "4.710989,-74.072090",
	},
	{
		LastName:        "Monty",
		FirstName:       "Jane",
		Email:           "driver2@example.com",
		CurrentLocation: "4.710789,-74.072290",
	},
}
