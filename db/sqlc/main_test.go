package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/emonoid/islami_bank_go_backend/utils"
	_ "github.com/lib/pq"
)


var testQueries *Queries
var testConn *sql.DB

func TestMain(m *testing.M){ 
    config, loadErr := utils.LoadConfig("../..")
	if loadErr != nil{
		log.Fatal("Cannot load config file", loadErr)
	}

	var err error
	testConn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}


	testQueries = New(testConn)

	    // Cleanup before test run
	testQueries.db.ExecContext(context.Background(), "DELETE FROM accounts")
	testQueries.db.ExecContext(context.Background(), "DELETE FROM entries")
	testQueries.db.ExecContext(context.Background(), "DELETE FROM transfers")
	
	os.Exit(m.Run())
}