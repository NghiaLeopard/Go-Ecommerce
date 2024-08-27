package initialize

import (
	"database/sql"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func initPostgresql() {
	sqlDB, err := sql.Open(global.Config.DBDrive, global.Config.DBSource)

	checkErrorPanic(err, "initialization postgresql err")

	global.Logger.Info("Initialization success")

	global.DB = db.New(sqlDB)

	setPool(sqlDB)
}

func setPool(db *sql.DB) {
	db.SetConnMaxIdleTime(time.Duration(global.Config.MaxIdleTime))
	db.SetMaxOpenConns(global.Config.MaxOpenConnect)
	db.SetConnMaxLifetime(time.Duration(global.Config.ConnMaxLifeTime))
}
