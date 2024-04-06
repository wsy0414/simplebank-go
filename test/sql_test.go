package test

import (
	"database/sql"
	"log"
	"os"
	"simplebank/db/sqlc"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver       = "postgres"
	driverSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var testQuery *sqlc.Queries
var db *sql.DB
var err error

func TestMain(m *testing.M) {
	db, err = sql.Open(driver, driverSource)
	if err != nil {
		log.Fatalln("connect db error")
	}

	testQuery = sqlc.New(db)

	os.Exit(m.Run())
}
