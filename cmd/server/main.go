package main

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/initialize"
)

// @title           Swagger Golang Ecommerce
// @description     This is a server ecommerce.

// @contact.name   Nguyễn Đại Nghĩa
// @contact.url    https://www.facebook.com/shy.leopard.beo
// @contact.email  nghiabeo1605@gmail.com

// @host      localhost:8080

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @SecurityDefinitions.apikey	BearerAuth
// @Name						Authorization
// @In							header
// @Description					Add prefix of Bearer before  token Ex: "Bearer token"

func main() {
	initialize.Run()
}
