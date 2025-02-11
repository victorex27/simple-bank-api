package main

import (
	"database/sql"
	"log"

	"github.com.victorex27/simple_bank/api"
	db "github.com.victorex27/simple_bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:1759@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = ":3000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
