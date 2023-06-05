package transactionDto

type TransactionResponse struct {
	ID      uint   `json:"id"`
	Status  string `json:"status"`
	WompiID string `json:"transaction_id"`
	Message string `json:"message"`
	Fare    uint   `json:"fare"`
}
