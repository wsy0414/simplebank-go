package db

import (
	"context"
	"database/sql"
	"log"
	"simplebank/config"
	"simplebank/db/sqlc"
	"sync"

	_ "github.com/lib/pq"
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
		c, err := config.LoadConfig()
		if err != nil {
			return
		}
		d, ee := sql.Open(c.Database.Driver, c.Database.Source)
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
