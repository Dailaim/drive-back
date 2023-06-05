package transaction

import (
	transactionServices "github.com/Daizaikun/drive-back/app/modules/transaction/services"
	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {

	app := fiber.New()

	app.Get("/transaction/wompi/:id", transactionServices.WompiTransactionByID)

	app.Get("/transaction/id/:id", transactionServices.TransactionByID)

	return app

}
