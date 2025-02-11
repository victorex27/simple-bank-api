package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com.victorex27/simple_bank/util"
	_ "github.com/lib/pq" // we are telling go formatter not to hide this package
)

var testQueries *Queries
var testDB *sql.DB

// 09:39:02 cannot connect to db: sql: unknown driver "postgres"
// to fix this error, we need to install the postgres driver
// go get github.com/lib/pq
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Failed to load env", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
