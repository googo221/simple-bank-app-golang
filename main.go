package main

import (
	"database/sql"
	"log"

	"github.com/golang-projects/simple-bank-app/api"
	db "github.com/golang-projects/simple-bank-app/db/sqlc"
	"github.com/golang-projects/simple-bank-app/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}