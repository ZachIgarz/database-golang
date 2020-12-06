package product

import "time"

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdateAt     time.Time
}
//Models slice of Model
type Models []*Model

type Storage interface {
	Migrate () error
	Create(*Model) error
	/*
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error*/
}

//Service of product
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
//Create is use for create a product
func (service *Service) Create(model *Model) error {
	model.CreatedAt = time.Now()
	return service.storage.Create(model)
}