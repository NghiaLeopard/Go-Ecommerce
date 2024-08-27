package initialize

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/logger"
)

func initLogger() {
	global.Logger = logger.NewLog()
}
