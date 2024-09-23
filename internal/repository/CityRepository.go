package repository

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type CityRepository struct{}

func NewCityRepository() IRepository.City {
	return &CityRepository{}
}

func (c *CityRepository) CreateCity(ctx *gin.Context, name string) (city db.City, err error) {
	city, err = global.DB.CreateCity(ctx, name)

	return
}

func (c *CityRepository) UpdateCity(ctx *gin.Context, id int64, name string) (city db.City, err error) {
	arg := db.UpdateCityParams{
		ID:   id,
		Name: name,
	}

	city, err = global.DB.UpdateCity(ctx, arg)

	return
}

func (c *CityRepository) GetCityById(ctx *gin.Context, id int64) (city db.City, err error) {
	city, err = global.DB.GetCityById(ctx, id)

	return
}

func (c *CityRepository) GetAllCity(ctx *gin.Context, page int, limit int, search string, order string) (city []db.City, err error) {

	offset := limit * (page - 1)

	arg := db.ListCityParams{
		Limit:   int32(limit),
		Offset:  int32(offset),
		OrderBy: order,
		Search:  search,
	}
	city, err = global.DB.ListCity(ctx, arg)

	return
}

func (c *CityRepository) GetCityByName(ctx *gin.Context, name string) (city db.City, err error) {
	city, err = global.DB.GetCityByName(ctx, name)

	return
}

func (c *CityRepository) DeleteCityById(ctx *gin.Context, id int64) (err error) {
	err = global.DB.DeleteCityById(ctx, id)

	return
}

func (c *CityRepository) DeleteManyCityByIds(ctx *gin.Context, arrayID []int64) (err error) {

	err = global.DB.DeleteManyCityByIds(ctx, arrayID)

	return
}
