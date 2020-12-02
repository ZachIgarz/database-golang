package main

import (
	"DataBases-Go/pkg/product"
	"DataBases-Go/storage"
	"log"
)

func main() {
	//instancia unica de la bd
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)


	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Product.Migrate : %v", err)
	}

}
