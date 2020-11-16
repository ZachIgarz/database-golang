package invoiceheader

import "time"

//Model of invoiceheader
type Model struct {
	ID uint
	Client string
	CreatedAt time.Time
	UpdateAt time.Time
}
