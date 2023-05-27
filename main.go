package main

import (
	"database/sql"
	"log"

	"github.com/jaspier/gobank/api"
	db "github.com/jaspier/gobank/db/sqlc"
	"github.com/jaspier/gobank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDdiver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
