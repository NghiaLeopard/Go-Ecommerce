package initialize

import (
	"log"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/wire"
	"go.uber.org/zap"
)

func Run() {
	initLoadConfig()
	initPostgresql()
	initLogger()
	initGmail()

	global.Logger.Info("Config success", zap.String("Status", "Success"))
	
	server, err := wire.InitServer(global.DB, global.Config)

	if err != nil {
		log.Fatal("Run server fail: ", err)
	}

	server.Start()
}
