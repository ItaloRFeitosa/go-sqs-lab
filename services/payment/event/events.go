package payment

import "time"

type OrderCreated struct {
	OrderID   string    `json:"orderId"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type InvoiceCreated struct {
	InvoiceID      string
	ConciliationID string
	Amount         int64
	QrCode         string
	CreatedAt      time.Time
}

type InvoicePaid struct {
	InvoiceID      string
	ConciliationID string
	AmountPaid     int64
	PaidAt         time.Time
}
