package event

import "time"

type OrderCreated struct {
	OrderID   string    `json:"orderId"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderPlaced struct {
	OrderID  string
	QrCode   string
	PlacedAt time.Time
}

type OrderPaid struct {
	OrderID    string
	AmountPaid int64
	PaidAt     time.Time
}
