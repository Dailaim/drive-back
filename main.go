package main

import (
	"github.com/Daizaikun/wompi-back/app/config"
	"github.com/Daizaikun/wompi-back/app/db"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func init() {
	db.Connect()
}

func main() {
	config.Run().Listen(":8080")
}
