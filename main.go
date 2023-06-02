package main

import (
	"github.com/Daizaikun/wompi-back/app/config"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()
}

func main() {
	config.Run().Listen(":8080")
}
