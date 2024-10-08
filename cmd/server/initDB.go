package main

import (
	"context"
	"database/sql"
	"log"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	_ "github.com/lib/pq"
)

func initDB() {

	arg1 := db.CreateRoleByDefaultParams{
		Name:       "Admin",
		Permission: []string{(constant.CONFIG_PERMISSIONS["ADMIN"].(string))},
	}

	arg2 := db.CreateRoleByDefaultParams{
		Name:       "Basic",
		Permission: []string{(constant.CONFIG_PERMISSIONS["BASIC"].(string))},
	}

	configEnv1, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Connect file env fail: ", err)
	}

	sqlDb, err := sql.Open(configEnv1.DBDrive, configEnv1.DBSource)

	if err != nil {
		log.Fatal("Connect to database fail: ", err)
	}

	sqlcDB := db.New(sqlDb)

	role, err := sqlcDB.CreateRoleByDefault(context.Background(), arg1)

	if err != nil {
		log.Fatal("create role fail: ", err)
	}

	_, err = sqlcDB.CreateRoleByDefault(context.Background(), arg2)

	if err != nil {
		log.Fatal("create role fail: ", err)
	}

	password, err := utils.HashPassword("1234567890@1nN")

	if err != nil {
		log.Fatal("hash password password fail: ", err)
	}

	arg := db.InitDefaultAdminParams{
		Email:    "admin@gmail.com",
		Password: password,
		Role:     sql.NullInt64{Int64: role.ID, Valid: role.ID != 0},
	}

	_, err = sqlcDB.InitDefaultAdmin(context.Background(), arg)

	if err != nil {
		log.Fatal("create user fail: ", err)
	}

}
