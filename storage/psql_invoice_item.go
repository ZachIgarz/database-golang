package storage
import (
"database/sql"
"fmt"
"github.com/ansel1/merry"
)

const psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
id SERIAL NOT NULL, 
invoice_header_id INT NOT NULL,
product_id INT NOT NULL,
create_at TIMESTAMP NOT NULL DEFAULT now(),
update_at TIMESTAMP,

CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT  ON DELETE RESTRICT,
CONSTRAINT invoice_items_product_id_fk FOREIGN KEY 
(product_id) REFERENCES products (id) ON UPDATE RESTRICT  ON DELETE RESTRICT 
)`

//psqlInvoiceItem used for work with postgres - product
type psqlInvoiceItem struct {
	db *sql.DB
}

//NewPsqlInvoiceItem return a new pointer of  psqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *psqlInvoiceItem {

	return &psqlInvoiceItem{db}
}

//Migrate implement the interface invoiceItem.Storage
func (p *psqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)

	if err != nil {

		return merry.New("error statement" + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {

		return merry.New("error statement exec" + err.Error())
	}
	fmt.Println("invoice item migration ok")

	return nil
}
