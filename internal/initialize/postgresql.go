package initialize

import (
	"database/sql"
	"log"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
)

func initPostgresql() {
	sqlDB, err := sql.Open(global.Config.DBDrive, global.Config.DBSource)

	if err != nil {
		log.Fatal("Connect to database fail: ", err)
	}

	global.DB = db.New(sqlDB)
}
