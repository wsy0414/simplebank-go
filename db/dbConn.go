package db

import (
	"context"
	"database/sql"
	"log"
	"simplebank/db/sqlc"
	"sync"

	_ "github.com/lib/pq"
)

const (
	driver       = "postgres"
	driverSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var db *sql.DB
var err error
var once sync.Once

type ConnData struct {
	Conn *sql.Conn
	Q    *sqlc.Queries
}

func GetQuery() (*ConnData, error) {
	once.Do(func() {
		log.Println("craete db connection!!!!!!")
		d, ee := sql.Open(driver, driverSource)
		if ee != nil {
			log.Print(ee)
			return
		}

		db = d
	})
	log.Println("get conn")
	log.Println("totoal already open connecion:", db.Stats().OpenConnections)
	log.Println("totoal is use connecion:", db.Stats().InUse)
	log.Println("totoal idle connecion:", db.Stats().Idle)
	conn, connErr := db.Conn(context.Background())
	if connErr != nil {
		return nil, connErr
	}

	return &ConnData{conn, sqlc.New(conn)}, nil
}
