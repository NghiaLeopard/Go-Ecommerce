package initialize

import (
	"log"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/wire"
	"go.uber.org/zap"
)

func Run() {
	initLoadConfig()
	initLogger()
	initPostgresql()
	initRedis()
	initGmail()
	initToken()

	global.Logger.Info("Config success", zap.String("Status", "Success"))

	server, err := wire.InitServer(global.DB, global.Config, global.Rdb)

	if err != nil {
		log.Fatal("Run server fail: ", err)
	}

	server.Start()
}
