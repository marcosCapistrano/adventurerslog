package db

import (
	"database/sql"
	"log"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "soil:password@/adventurerslog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
