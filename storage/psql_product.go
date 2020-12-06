package storage

import (
	"DataBases-Go/pkg/product"
	"database/sql"
	"fmt"
	"github.com/ansel1/merry"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
id SERIAL NOT NULL, 
name VARCHAR(25) NOT NULL,
observations VARCHAR(100),
price INT NOT NULL,
create_at TIMESTAMP NOT NULL DEFAULT now(),
update_at TIMESTAMP,
CONSTRAINT products_id_pk PRIMARY KEY (id))`

psqlCreateProduct = `INSERT INTO products(name, observations, price, create_at) VALUES($1, $2, $3, $4) RETURNING id`
)

//PsqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

//NewPsqlProduct return a new pointer of  PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {

	return &PsqlProduct{db}
}

//Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)

	if err != nil {

		return merry.New("error statement" + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {

		return merry.New("error statement exec" + err.Error())
	}

	return nil
}

//Create implement the interface of product storage
func (p *PsqlProduct) Create (model *product.Model) error {

	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {

		return merry.New("error prepare create of product")
	}
	defer stmt.Close()

	err = stmt.QueryRow(model.Name,
		model.Observations,
		model.Price,
		model.CreatedAt,
	).Scan(&model.ID)

	if err != nil {
		return merry.New("error scan query product")
	}
	fmt.Println("product create!")

	return nil
}