package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
