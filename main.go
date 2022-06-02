package main

import (
	"database/sql"
	"log"

	api "github.com/galib612/simplebank/api"
	db "github.com/galib612/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serveraddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serveraddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
