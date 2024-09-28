package repository

import (
	"database/sql"
	"sync"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func (r *ProductRepository) GetProductBySlug(ctx *gin.Context, slug string, isViewed bool) (db.GetProductBySlugRow, error) {

	Product, err := global.DB.GetProductBySlug(ctx, slug)

	if err == nil && isViewed == true {
		var wg sync.WaitGroup
		wg.Add(1)
		view := Product.Views.Int32 + 1
		go r.UpdateViewProduct(ctx, Product.ID, view, &wg)

		wg.Wait()
	}

	return Product, err
}

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

func (r *ProductRepository) UpdateViewProduct(ctx *gin.Context, id int64, view int32, wg *sync.WaitGroup) {
	arg := db.UpdateViewProductParams{
		ID:    id,
		Views: sql.NullInt32{Int32: view, Valid: true},
	}
	err := global.DB.UpdateViewProduct(ctx, arg)

	if err == nil {
		wg.Done()
	} else {
		global.Logger.Error(err.Error(), zap.String("status", "Error"))
	}
}

func (r *ProductRepository) UpdateUniqueView(ctx *gin.Context, productId int64, wg *sync.WaitGroup) {
	payload := ctx.MustGet(constant.AuthorizationKey).(*token.Payload)
	arg := db.CreateProductUniqueViewParams{
		ProductID: int32(productId),
		UserID:    int32(payload.Id),
	}

	err := global.DB.CreateProductUniqueView(ctx, arg)

	if err == nil {
		wg.Done()
	} else {
		global.Logger.Error(err.Error(), zap.String("status", "Error"))
	}
}

func (r *ProductRepository) DeleteProduct(ctx *gin.Context, id int64) error {
	err := global.DB.DeleteProductById(ctx, id)

	return err
}

func (r *ProductRepository) DeleteManyProduct(ctx *gin.Context, arrayId []int64) error {
	err := global.DB.DeleteManyProductsByIds(ctx, arrayId)

	return err
}
