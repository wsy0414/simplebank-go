package sqlc

import (
	"database/sql"
	"os"
	"simplebank/config"
	_ "simplebank/flags"
	"testing"

	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config.LoadConfig("../../config")
	db, err := sql.Open(config.ConfigVal.Database.Driver, config.ConfigVal.Database.Source)
	if err != nil {
		panic(err.Error())
	}

	testStore = NewStore(db)
	os.Exit(m.Run())
}
