package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
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

// CreateCityUseCase implements IUseCase.ICityUseCase.
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

// GetCityById implements IUseCase.ICityUseCase.
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
