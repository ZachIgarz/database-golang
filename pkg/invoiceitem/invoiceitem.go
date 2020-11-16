package invoiceitem

import "time"
//Model of invoiceitem
type Model struct {
	ID uint
	InvoiceHeaderID uint
	ProductID uint
	CreatedAt time.Time
	UpdateAt time.Time
}
