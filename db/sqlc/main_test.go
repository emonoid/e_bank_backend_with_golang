package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
 

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5454/islami_bank?sslmode=disable"
)

var testQueries *Queries
var testConn *sql.DB

func TestMain(m *testing.M){ 
	var err error
	testConn, err = sql.Open(dbDriver, dbSource)
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