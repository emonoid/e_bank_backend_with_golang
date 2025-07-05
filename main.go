package main

import (
	"database/sql"
	"log"

	"github.com/emonoid/islami_bank_go_backend/api"
	db "github.com/emonoid/islami_bank_go_backend/db/sqlc"
	"github.com/emonoid/islami_bank_go_backend/utils"
	_ "github.com/lib/pq"
)



func main(){
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start servcer:", err)
	}
}
