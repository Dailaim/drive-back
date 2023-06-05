package transactionDto

type ErrorResponse struct {
	Error Message `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}
