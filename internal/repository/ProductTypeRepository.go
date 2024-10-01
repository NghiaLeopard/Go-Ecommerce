package repository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type ProductTypeRepository struct{}

func NewProductTypeRepository() IRepository.ProductType {
	return &ProductTypeRepository{}
}

func (r *ProductTypeRepository) CreateProductType(ctx *gin.Context, name string, slug string) (ProductType db.ProductType, err error) {
	arg := db.CreateProductTypeParams{
		Name: name,
		Slug: slug,
	}

	ProductType, err = global.DB.CreateProductType(ctx, arg)

	return
}

func (r *ProductTypeRepository) GetProductTypeById(ctx *gin.Context, id int64) (db.ProductType, error) {
	ProductType, err := global.DB.GetProductTypeById(ctx, id)

	return ProductType, err
}

func (r *ProductTypeRepository) GetProductTypeByName(ctx *gin.Context, name string) (db.ProductType, error) {
	ProductType, err := global.DB.GetProductTypeByName(ctx, name)

	return ProductType, err
}

func (r *ProductTypeRepository) GetAllProductType(ctx *gin.Context, req IRequest.GetAllProductType) ([]db.ListProductTypeRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.ListProductTypeParams{
		Limit:   req.Limit,
		Offset:  offset,
		Search:  req.Search,
		OrderBy: req.Order,
	}

	ProductType, err := global.DB.ListProductType(ctx, arg)

	return ProductType, err
}

func (r *ProductTypeRepository) UpdateProductType(ctx *gin.Context, id int64, name string, slug string) (db.ProductType, error) {
	arg := db.UpdateProductTypeParams{
		ID:   id,
		Name: name,
		Slug: slug,
	}
	ProductType, err := global.DB.UpdateProductType(ctx, arg)

	return ProductType, err
}

func (r *ProductTypeRepository) DeleteProductType(ctx *gin.Context, id int64) error {
	err := global.DB.DeleteProductTypeById(ctx, id)

	return err
}

func (r *ProductTypeRepository) DeleteManyProductType(ctx *gin.Context, arrayId []int64) error {
	err := global.DB.DeleteManyProductTypesByIds(ctx, arrayId)

	return err
}
