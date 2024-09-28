package repository

import (
	"database/sql"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type ProductRepository struct{}

func NewProductRepository() IRepository.Product {
	return &ProductRepository{}
}

func (r *ProductRepository) CreateProductNotDiscount(ctx *gin.Context, req IRequest.CreateProduct) (Product db.Product, err error) {
	arg := db.CreateProductParams{
		Name:         req.Name,
		Slug:         req.Slug,
		Description:  req.Description,
		Type:         req.Type,
		Price:        req.Price,
		CountInStock: req.CountInStock,
		Status:       req.Status,
		Image:        req.Image,
		Location:     req.Location,
	}

	Product, err = global.DB.CreateProduct(ctx, arg)

	return
}

func (r *ProductRepository) CreateProductDiscount(ctx *gin.Context, req IRequest.CreateProduct) (Product db.Product, err error) {
	arg := db.CreateProductParams{
		Name:              req.Name,
		Slug:              req.Slug,
		Description:       req.Description,
		Type:              req.Type,
		Price:             req.Price,
		CountInStock:      req.CountInStock,
		Status:            req.Status,
		Image:             req.Image,
		Location:          req.Location,
		Discount:          sql.NullInt32{Int32: req.Discount, Valid: true},
		DiscountStartDate: sql.NullTime{Time: req.DiscountStart, Valid: true},
		DiscountEndDate:   sql.NullTime{Time: req.DiscountEndDate, Valid: true},
	}

	Product, err = global.DB.CreateProduct(ctx, arg)

	return
}

func (r *ProductRepository) GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error) {
	Product, err := global.DB.GetProductById(ctx, id)
	return Product, err
}

// func (r *ProductRepository) GetProductByName(ctx *gin.Context, name string) (db.Product, error) {
// 	Product, err := global.DB.GetProductByName(ctx, name)

// 	return Product, err
// }

// func (r *ProductRepository) GetAllProduct(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error) {

// 	offset := req.Limit * (req.Page - 1)
// 	arg := db.ListProductParams{
// 		Limit:   req.Limit,
// 		Offset:  offset,
// 		Search:  req.Search,
// 		OrderBy: req.Order,
// 	}

// 	Product, err := global.DB.ListProduct(ctx, arg)

// 	return Product, err
// }

// func (r *ProductRepository) UpdateProduct(ctx *gin.Context, id int64, name string, slug string) (db.Product, error) {
// 	arg := db.UpdateProductParams{
// 		ID:   id,
// 		Name: name,
// 		Slug: slug,
// 	}
// 	Product, err := global.DB.UpdateProduct(ctx, arg)

// 	return Product, err
// }

func (r *ProductRepository) DeleteProduct(ctx *gin.Context, id int64) error {
	err := global.DB.DeleteProductById(ctx, id)

	return err
}

func (r *ProductRepository) DeleteManyProduct(ctx *gin.Context, arrayId []int64) error {
	err := global.DB.DeleteManyProductsByIds(ctx, arrayId)

	return err
}
