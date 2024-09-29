package usecase

import (
	"fmt"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductTypeUseCase struct {
	ProductTypeRepo IRepository.ProductType
}

func NewProductTypeUseCase(productTypeRepo IRepository.ProductType) IUseCase.ProductType {
	return &ProductTypeUseCase{
		ProductTypeRepo: productTypeRepo,
	}
}

func (c *ProductTypeUseCase) CreateProductType(ctx *gin.Context, name string, slug string) (IResponse.ProductType, error, int) {
	_, err := c.ProductTypeRepo.GetProductTypeByName(ctx, name)

	if err == nil {
		global.Logger.Error("ProductType is exist", zap.String("Status", "Error"))
		return IResponse.ProductType{}, fmt.Errorf("ProductType is exist"), 409
	}

	ProductType, err := c.ProductTypeRepo.CreateProductType(ctx, name, slug)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.ProductType{}, err, 401
	}

	return IResponse.ProductType{
		Id:       ProductType.ID,
		Name:     ProductType.Name,
		Slug:     ProductType.Slug,
		CreateAt: ProductType.CreateAt,
		UpdateAt: ProductType.UpdateAt.Time,
	}, nil, 201
}

func (c *ProductTypeUseCase) GetAllProductTypeUseCase(ctx *gin.Context, req IRequest.GetAllProductType) ([]db.ProductType, error, int) {
	ProductType, err := c.ProductTypeRepo.GetAllProductType(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return []db.ProductType{}, fmt.Errorf("get ProductType is not exist"), 401
	}

	return ProductType, nil, 200
}

func (c *ProductTypeUseCase) GetProductTypeUseCase(ctx *gin.Context, id int) (IResponse.ProductType, error, int) {
	ProductType, err := c.ProductTypeRepo.GetProductTypeById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.ProductType{}, fmt.Errorf("get ProductType is not exist"), 401
	}

	return IResponse.ProductType{
		Id:       ProductType.ID,
		Name:     ProductType.Name,
		Slug:     ProductType.Slug,
		CreateAt: ProductType.CreateAt,
		UpdateAt: ProductType.UpdateAt.Time,
	}, nil, 200
}

func (c *ProductTypeUseCase) UpdateProductTypeUseCase(ctx *gin.Context, id int, name string, slug string) (IResponse.ProductType, error, int) {
	idInt64 := int64(id)

	ProductType, err := global.DB.GetProductTypeById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.ProductType{}, fmt.Errorf("ProductType is not exist"), 401
	}

	ProductType, err = c.ProductTypeRepo.UpdateProductType(ctx, idInt64, name, slug)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.ProductType{}, fmt.Errorf("update ProductType is fail"), 401
	}

	res := IResponse.ProductType{
		Id:       ProductType.ID,
		Name:     ProductType.Name,
		Slug:     ProductType.Slug,
		CreateAt: ProductType.CreateAt,
		UpdateAt: ProductType.UpdateAt.Time,
	}

	return res, nil, 200
}

func (c *ProductTypeUseCase) DeleteProductTypeUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeleteProductTypeById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get ProductType is not exist"), 401
	}

	return nil, 200
}

func (c *ProductTypeUseCase) DeleteManyProductTypeUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayId64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayId64[i] = int64(v)
	}

	err := global.DB.DeleteManyProductTypesByIds(ctx, arrayId64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get ProductType is not exist"), 401
	}

	return nil, 200
}
