package initialize

import (
	"context"
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func initRedis() {
	r := global.Config
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.HostRedis, r.PortRedis),
		Password: r.PasswordRedis,
		DB:       r.DbRedis,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		checkErrorPanic(err, "initialization redis err")
	}

	global.Logger.Info("init redis success")
	global.Rdb = rdb
}
