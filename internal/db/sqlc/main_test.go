package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDb *sql.DB



func TestMain(m *testing.M) {
	config, err := config.LoadConfig("..//..//..")

	if err != nil {
		log.Fatal("Connect fail: ",err)
	}

	testDb,err = sql.Open(config.DBDrive,config.DBSource)

	if err != nil {
		log.Fatal("Connect to database fail: ",err)
	}

	testQuery = New(testDb)

	os.Exit(m.Run())

}

