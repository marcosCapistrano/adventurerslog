package main

import (
	"db"
)

func main() {
	db := db.Open()

	server := NewServer(db)
	server.Query()
}
