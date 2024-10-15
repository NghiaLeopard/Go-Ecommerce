package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeliveryUseCase struct {
	DeliveryRepo IRepository.Delivery
}

func NewDeliveryUseCase(DeliveryRepo IRepository.Delivery) IUseCase.Delivery {
	return &DeliveryUseCase{DeliveryRepo: DeliveryRepo}
}

func (c *DeliveryUseCase) CreateDeliveryUseCase(ctx *gin.Context, req IRequest.CreateDelivery) (IResponse.Delivery, error, int) {
	_, err := c.DeliveryRepo.GetDeliveryByName(ctx, req.Name)

	if err == nil {
		global.Logger.Error("delivery is  exist", zap.String("Status", "Error"))
		return IResponse.Delivery{}, fmt.Errorf("delivery is  exist"), 409
	}

	Delivery, err := c.DeliveryRepo.CreateDelivery(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Delivery{}, err, 401
	}

	return IResponse.Delivery{
		Id:    Delivery.ID,
		Name:  Delivery.Name,
		Price: Delivery.Price,
	}, nil, 201
}

func (c *DeliveryUseCase) GetDeliveryUseCase(ctx *gin.Context, id int) (IResponse.Delivery, error, int) {

	Delivery, err := c.DeliveryRepo.GetDeliveryById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Delivery{}, fmt.Errorf("get Delivery is not exist"), 401
	}

	return IResponse.Delivery{
		Id:       Delivery.ID,
		Name:     Delivery.Name,
		Price:    Delivery.Price,
		CreateAt: Delivery.CreateAt,
	}, nil, 200
}

func (c *DeliveryUseCase) GetAllDeliveryUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllDelivery, error, int) {
	Delivery, err := c.DeliveryRepo.GetAllDelivery(ctx, page, limit, search, order)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllDelivery{}, fmt.Errorf("get Delivery is not exist"), 401
	}

	fmt.Println(Delivery, limit, order)

	if len(Delivery) == 0 {
		return IResponse.GetAllDelivery{
			DeliveryTypes: Delivery,
			TotalCount:    0,
			TotalPage:     0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(limit), Delivery[0].TotalCount)

	return IResponse.GetAllDelivery{
		DeliveryTypes: Delivery,
		TotalCount:    Delivery[0].TotalCount,
		TotalPage:     totalPage,
	}, nil, 200
}

func (c *DeliveryUseCase) UpdateDeliveryUseCase(ctx *gin.Context, id int, body IRequest.GetBodyUpdateDelivery) (IResponse.Delivery, error, int) {
	idInt64 := int64(id)

	_, err := global.DB.GetDeliveryById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Delivery{}, fmt.Errorf("Delivery is not exist"), 401
	}

	Delivery, err := c.DeliveryRepo.UpdateDelivery(ctx, idInt64, body)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Delivery{}, fmt.Errorf("update Delivery is fail"), 401
	}

	res := IResponse.Delivery{
		Id:    Delivery.ID,
		Name:  Delivery.Name,
		Price: Delivery.Price,
	}

	return res, nil, 200
}

func (c *DeliveryUseCase) DeleteDeliveryUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeleteDeliveryById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Delivery is not exist"), 401
	}

	return nil, 200
}

func (c *DeliveryUseCase) DeleteManyDeliveryUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayInt64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayInt64[i] = int64(v)
	}

	err := global.DB.DeleteManyDeliveryByIds(ctx, arrayInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Delivery is not exist"), 401
	}

	return nil, 200
}
