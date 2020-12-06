package main

import (
	"DataBases-Go/pkg/product"
	"DataBases-Go/storage"
	"fmt"
)

func main() {
	//unique instance of fb
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	model := &product.Model{
		Name: "Curse of go",
		Price: 50,
	}
	err := serviceProduct.Create(model)
	if err != nil {

		fmt.Println("Error " + err.Error())
	}

	fmt.Printf("%+v\n", model)


}
