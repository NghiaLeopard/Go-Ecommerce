package main

import (
	"database/sql"
	"log"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/wire"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Connect file env fail: ", err)
	}

	sqlDB, err := sql.Open(config.DBDrive, config.DBSource)

	if err != nil {
		log.Fatal("Connect to database fail: ", err)
	}

	sqlcDB := db.New(sqlDB)

	Token, err := token.NewPasetoMaker([]byte(config.Symmetric))

	if err != nil {
		log.Fatal("Init token fail: ", err)
	}

	server, err := wire.InitApi(sqlcDB, config, Token)

	if err != nil {
		log.Fatal("Run server fail: ", err)
	}

	server.Start()
}
