package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	//This function will be executed only once, even if it is called several times
	once.Do(func() {
		var connectionURL = "postgres://postgres:Abcd-456@localhost/go-db?sslmode=disable"
		db, err := sql.Open("postgres", connectionURL)
		if err != nil {
		log.Fatalf("can't open db : %v", err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping : %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}
//Pool return a unique instace of db
func Pool() *sql.DB{
	return db
}
