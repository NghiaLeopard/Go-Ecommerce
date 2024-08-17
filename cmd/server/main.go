package main

import (
	"database/sql"
	"log"

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

	db, err := sql.Open(config.DBDrive, config.DBSource)

	if err != nil {
		log.Fatal("Connect to database fail: ", err)
	}

	Token, err := token.NewPasetoMaker([]byte(config.Symmetric))

	if err != nil {
		log.Fatal("Init token fail: ", err)
	}

	server, err := wire.InitApi(db, config, Token)

	if err != nil {
		log.Fatal("Run server fail: ", err)
	}

	server.Start()
}
