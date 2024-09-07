package repository

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type CityRepository struct{}

func NewCityRepository() IRepository.ICity {
	return &CityRepository{}
}

// CreateCity implements IRepository.ICity
func (c *CityRepository) CreateCity(ctx *gin.Context, name string) (city db.City, err error) {
	city, err = global.DB.CreateCity(ctx, name)

	return
}
