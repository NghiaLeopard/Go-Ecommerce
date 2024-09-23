package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
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

func (c *CityUseCase) CreateCityUseCase(ctx *gin.Context, name string) (IResponse.City, error, int) {
	_, err := c.CityRepo.GetCityByName(ctx, name)

	if err == nil {
		global.Logger.Error("city is  exist", zap.String("Status", "Error"))
		return IResponse.City{}, fmt.Errorf("city is  exist"), 409
	}

	city, err := global.DB.CreateCity(ctx, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.City{}, err, 401
	}

	global.Logger.Info("Create city", zap.String("Status", "Error"))
	return IResponse.City{
		Id:   city.ID,
		Name: city.Name,
	}, nil, 201
}

func (c *CityUseCase) GetCityUseCase(ctx *gin.Context, id int) (IResponse.City, error, int) {

	city, err := c.CityRepo.GetCityById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.City{}, fmt.Errorf("get city is not exist"), 401
	}

	return IResponse.City{
		Id:       city.ID,
		Name:     city.Name,
		CreateAt: city.CreateAt,
	}, nil, 200
}

func (c *CityUseCase) GetAllCityUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) ([]db.City, error, int) {

	city, err := c.CityRepo.GetAllCity(ctx, page, limit, search, order)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return []db.City{}, fmt.Errorf("get city is not exist"), 401
	}

	return city, nil, 200
}

func (c *CityUseCase) UpdateCityUseCase(ctx *gin.Context, id int, name string) (IResponse.City, error, int) {
	idInt64 := int64(id)

	_, err := global.DB.GetCityById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.City{}, fmt.Errorf("city is not exist"), 401
	}

	city, err := c.CityRepo.UpdateCity(ctx, idInt64, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.City{}, fmt.Errorf("update city is fail"), 401
	}

	res := IResponse.City{
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
