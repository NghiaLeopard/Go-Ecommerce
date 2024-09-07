package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
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
		return response.ICityResponse{}, fmt.Errorf("city is exist"), 409
	}

	city, err := global.DB.CreateCity(ctx, name)

	if err != nil {
		return response.ICityResponse{}, err, 401
	}

	return response.ICityResponse{
		Id:   int(city.ID),
		Name: city.Name,
	}, nil, 201
}
