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


type Storage interface {
	Migrate () error

}

//Service of  invoiceitem
type Service struct {
	storage Storage
}
//NewService return a pointer of service
func NewService(s Storage) *Service{

	return &Service{s}
}

//Migrate is use for migrate product
func (service *Service) Migrate()  error{

	return service.storage.Migrate()
}