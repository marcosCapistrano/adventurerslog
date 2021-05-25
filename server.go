package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db}
}

func (s *Server) Query() {
	var name string
	err := s.db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
