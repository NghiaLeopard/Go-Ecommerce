package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Nguyen Dai Nghia Dep trai")

	config,err := config.LoadConfig(".")

	r := gin.Default()


	r.GET("/user",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message": "Nghia dep trai",
		})
	})

	r.Run("0.0.0.0:8000")

	if err != nil {
		log.Fatal("Connect file env fail: ",err)
	}

	fmt.Print(config.DBDrive)
}