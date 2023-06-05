package main

import (
	"fmt"

	"github.com/Daizaikun/drive-back/app/config"
	"github.com/Daizaikun/drive-back/app/lib"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

// @title Diver Back API
// @version 1.0
// @description The driver-back API description for swagger documentation
// @contact.name Leonardo Iglesias
// @contact.email laiglesias.min@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	app := config.Run()

	port := lib.CheckEnv("PORT", "8080")

	app.Listen(fmt.Sprintf(":%s", port))
}
