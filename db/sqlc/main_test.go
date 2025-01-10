package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // we are telling go formatter not to hide this package
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1759@localhost:5433/simple_bank?sslmode=disable"
)

// 09:39:02 cannot connect to db: sql: unknown driver "postgres"
// to fix this error, we need to install the postgres driver
// go get github.com/lib/pq
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
