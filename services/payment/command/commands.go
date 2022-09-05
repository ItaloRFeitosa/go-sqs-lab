package command

type CreateInvoice struct {
	ConciliationID string
	Amount         int64
}

type PayInvoice struct {
	InvoiceID string
	Amount    int64
}
