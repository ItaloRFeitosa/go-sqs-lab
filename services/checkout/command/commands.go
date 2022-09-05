package command

type PayOrder struct {
	OrderID string
	Amount  int64
}

type PlaceOrder struct {
	OrderID string
	QrCode  string
	BrCode  string
}
