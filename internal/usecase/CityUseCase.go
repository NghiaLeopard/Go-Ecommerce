package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CityUseCase struct {
	CityRepo IRepository.ICity
}

func NewCityUseCase(cityRepo IRepository.ICity) IUseCase.ICityUseCase {
	return &CityUseCase{CityRepo: cityRepo}
}

func (c *CityUseCase) CreateCityUseCase(ctx *gin.Context, name string) (response.ICityResponse, error, int) {
	_, err := global.DB.GetCityByName(ctx, name)

	if err == nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, fmt.Errorf("city is exist"), 409
	}

	city, err := global.DB.CreateCity(ctx, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, err, 401
	}

	global.Logger.Info("Create city", zap.String("Status", "Error"))
	return response.ICityResponse{
		Id:   city.ID,
		Name: city.Name,
	}, nil, 201
}

func (c *CityUseCase) GetCityUseCase(ctx *gin.Context, id int) (response.ICityResponse, error, int) {
	city, err := global.DB.GetCityById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, err, 401
	}

	return response.ICityResponse{
		Id:   city.ID,
		Name: city.Name,
	}, nil, 200
}

func (c *CityUseCase) UpdateCityUseCase(ctx *gin.Context, id int, name string) (response.ICityResponse, error, int) {
	city, err := global.DB.GetCityById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, err, 401
	}

	arg := db.UpdateCityParams{
		ID:   int64(id),
		Name: name,
	}
	city, err = global.DB.UpdateCity(ctx, arg)

	return response.ICityResponse{
		Id:   city.ID,
		Name: city.Name,
	}, nil, 200
}
