package main

import (
	"database/sql"
	"simplebank/api"
	"simplebank/config"
	"simplebank/db/sqlc"
	_ "simplebank/flags"

	_ "simplebank/docs"

	_ "github.com/lib/pq"
)

//	@title			Swagger SimpleBank API
//	@version		1.0
//	@description	This is a simple bank server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	config.LoadConfig("./config")
	db := connDB()
	server := api.NewServer(sqlc.NewStore(db))

	server.Run(config.ConfigVal.Server.Port)
}

// connDB return a sql package's DB implement
func connDB() *sql.DB {
	db, err := sql.Open(config.ConfigVal.Database.Driver, config.ConfigVal.Database.Source)
	if err != nil {
		panic(err.Error())
	}

	return db
}
