package main

import (
	"github.com/Daizaikun/drive-back/app/config"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app := config.Run()

	app.Listen(":8080")
}
