package wompiEventDto

import "time"

type Transaction struct {
	ID                string `json:"id"`
	AmountInCents     int    `json:"amount_in_cents"`
	Reference         string `json:"reference"`
	CustomerEmail     string `json:"customer_email"`
	Currency          string `json:"currency"`
	PaymentMethodType string `json:"payment_method_type"`
	RedirectURL       string `json:"redirect_url"`
	Status            string `json:"status"`
	ShippingAddress   string `json:"shipping_address"`
	PaymentLinkID     string `json:"payment_link_id"`
	PaymentSourceID   string `json:"payment_source_id"`
}

type Signature struct {
	Properties []string `json:"properties"`
	Checksum   string   `json:"checksum"`
}

type Event struct {
	Event       string    `json:"event"`
	Data        Data      `json:"data"`
	Environment string    `json:"environment"`
	Signature   Signature `json:"signature"`
	Timestamp   int64     `json:"timestamp"`
	SentAt      time.Time `json:"sent_at"`
}

type Data struct {
	Transaction Transaction `json:"transaction"`
}
