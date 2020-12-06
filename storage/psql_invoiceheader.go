package storage

import (
	"database/sql"
	"fmt"
	"github.com/ansel1/merry"
)

const psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
id SERIAL NOT NULL, 
client VARCHAR(100) NOT NULL,
create_at TIMESTAMP NOT NULL DEFAULT now(),
update_at TIMESTAMP,
CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id))`

//psqlInvoiceItem used for work with postgres - product
type PsqlInvoiceHeader struct {
	db *sql.DB
}

//NewPsqlInvoiceHeader return a new pointer of  psqlInvoiceItem
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {

	return &PsqlInvoiceHeader{db}
}

//Migrate implement the interface product.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)

	if err != nil {

		return merry.New("error statement" + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {

		return merry.New("error statement exec" + err.Error())
	}
	fmt.Println("invoice header migration ok")
	return nil
}
