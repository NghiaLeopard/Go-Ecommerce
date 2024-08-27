package initialize

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/gmail"
)

func initGmail() {
	global.Gmail = gmail.NewEmailSender(global.Config)
	global.Logger.Info("Initialization gmail success")
}
