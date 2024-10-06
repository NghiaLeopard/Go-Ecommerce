package initialize

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
)

func initToken() {
	init, err := token.NewJWTMaker(global.Config.SecretKey)

	if err != nil {
		checkErrorPanic(err, "Initialization token false")
	}

	global.Token = init
}
