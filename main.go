package main

import (
	"DataBases-Go/pkg/product"
	"DataBases-Go/storage"
	"fmt"
	"log"
)

func main() {
	//unique instance of fb
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("Product.Migrate : %v", err)
	}
	fmt.Println("Se realizo la migracion")

}
