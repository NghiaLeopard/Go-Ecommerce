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
	CityRepo IRepository.City
}

func NewCityUseCase(cityRepo IRepository.City) IUseCase.City {
	return &CityUseCase{CityRepo: cityRepo}
}

func (c *CityUseCase) CreateCityUseCase(ctx *gin.Context, name string) (response.ICityResponse, error, int) {
	_, err := c.CityRepo.GetCityByName(ctx, name)

	if err == nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, fmt.Errorf("city is not exist"), 409
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
	city, err := c.CityRepo.GetCityById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, fmt.Errorf("get city is not exist"), 401
	}

	return response.ICityResponse{
		Id:   city.ID,
		Name: city.Name,
	}, nil, 200
}

func (c *CityUseCase) UpdateCityUseCase(ctx *gin.Context, id int, name string) (response.ICityResponse, error, int) {
	idInt64 := int64(id)

	_, err := global.DB.GetCityById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, fmt.Errorf("city is not exist"), 401
	}

	city, err := c.CityRepo.UpdateCity(ctx, idInt64, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ICityResponse{}, fmt.Errorf("update city is fail"), 401
	}

	res := response.ICityResponse{
		Id:   city.ID,
		Name: city.Name,
	}

	return res, nil, 200
}

func (c *CityUseCase) DeleteCityUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeleteCityById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get city is not exist"), 401
	}

	return nil, 200
}

func (c *CityUseCase) DeleteManyCityUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayInt64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayInt64[i] = int64(v)
	}

	err := global.DB.DeleteManyCityByIds(ctx, arrayInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get city is not exist"), 401
	}

	return nil, 200
}
