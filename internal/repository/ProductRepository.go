package repository

import (
	"sync"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
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
		Discount:          req.Discount,
		DiscountStartDate: req.DiscountStart,
		DiscountEndDate:   req.DiscountEndDate,
	}

	Product, err = global.DB.CreateProduct(ctx, arg)

	return
}

func (r *ProductRepository) GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error) {
	Product, err := global.DB.GetProductById(ctx, id)
	return Product, err
}

func (r *ProductRepository) GetProductTypeBySlug(ctx *gin.Context, slug string) (db.GetProductTypeBySlugRow, error) {
	Product, err := global.DB.GetProductTypeBySlug(ctx, slug)
	return Product, err
}

func (r *ProductRepository) GetAllProductRelated(ctx *gin.Context, req IRequest.GetAllProductRelated, id int64, city int32) ([]db.GetAllProductRelatedRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.GetAllProductRelatedParams{
		Limit:  req.Limit,
		Offset: offset,
		ID:     id,
		Type:   city,
	}

	Product, err := global.DB.GetAllProductRelated(ctx, arg)
	return Product, err
}

func (r *ProductRepository) GetProductBySlug(ctx *gin.Context, slug string) (db.GetProductBySlugRow, error) {
	Product, err := global.DB.GetProductBySlug(ctx, slug)

	return Product, err
}

func (r *ProductRepository) GetProductPublicById(ctx *gin.Context, productId int64) (db.GetProductPublicByIdRow, error) {
	Product, err := global.DB.GetProductPublicById(ctx, productId)

	return Product, err
}

func (r *ProductRepository) GetAllProductAdmin(ctx *gin.Context, req IRequest.GetAllProductAdmin) ([]db.GetAllProductAdminRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.GetAllProductAdminParams{
		Limit:   req.Limit,
		Offset:  offset,
		Search:  req.Search,
		Status:  req.Status,
		Type:    req.ProductType,
		OrderBy: req.Order,
	}

	Product, err := global.DB.GetAllProductAdmin(ctx, arg)

	return Product, err
}

func (r *ProductRepository) GetAllProductPublic(ctx *gin.Context, req IRequest.GetAllProductPublic) ([]db.GetAllProductPublicRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.GetAllProductPublicParams{
		Limit:  req.Limit,
		Offset: offset,
		Search: req.Search,
		Status: req.Status,
		Type:   req.ProductType,
	}

	Product, err := global.DB.GetAllProductPublic(ctx, arg)

	return Product, err
}

func (r *ProductRepository) GetAllProductMeLiked(ctx *gin.Context, req IRequest.GetAllProductLiked, userId int) ([]db.GetAllProductLikeRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.GetAllProductLikeParams{
		Limit:  req.Limit,
		Offset: offset,
		Search: req.Search,
		UserID: int32(userId),
	}

	Product, err := global.DB.GetAllProductLike(ctx, arg)

	return Product, err
}

func (r *ProductRepository) GetAllProductMeViewed(ctx *gin.Context, req IRequest.GetAllProductViewed, userId int) ([]db.GetAllProductViewRow, error) {

	offset := req.Limit * (req.Page - 1)
	arg := db.GetAllProductViewParams{
		Limit:  req.Limit,
		Offset: offset,
		Search: req.Search,
		UserID: int32(userId),
	}

	Product, err := global.DB.GetAllProductView(ctx, arg)

	return Product, err
}

func (r *ProductRepository) UpdateProduct(ctx *gin.Context, id int64, body IRequest.UpdateProduct) (db.Product, error) {
	arg := db.UpdateProductParams{
		ID:                id,
		Name:              body.Name,
		Slug:              body.Slug,
		Price:             body.Price,
		CountInStock:      body.CountInStock,
		Description:       body.Description,
		Discount:          body.Discount,
		DiscountStartDate: body.DiscountStart,
		DiscountEndDate:   body.DiscountEndDate,
		Type:              body.Type,
		Location:          body.Location,
		Status:            body.Status,
		Image:             body.Image,
	}

	Product, err := global.DB.UpdateProduct(ctx, arg)

	return Product, err
}

func (r *ProductRepository) UpdateViewProduct(ctx *gin.Context, id int64, view int32, wg *sync.WaitGroup) {
	arg := db.UpdateViewProductParams{
		ID:    id,
		Views: view,
	}
	err := global.DB.UpdateViewProduct(ctx, arg)

	if err == nil {
		wg.Done()
	} else {
		global.Logger.Error(err.Error(), zap.String("status", "Error"))
	}
}

func (r *ProductRepository) UpdateUniqueView(ctx *gin.Context, productId int64, userId int, wg *sync.WaitGroup) {
	arg := db.CreateProductUniqueViewParams{
		ProductID: int32(productId),
		UserID:    int32(userId),
	}

	err := global.DB.CreateProductUniqueView(ctx, arg)

	if err == nil {
		wg.Done()
	} else {
		global.Logger.Error(err.Error(), zap.String("status", "Error"))
	}
}

func (r *ProductRepository) UpdateLikeProduct(ctx *gin.Context, productId int64, userId int) error {
	arg := db.CreateProductLikeParams{
		ProductID: int32(productId),
		UserID:    int32(userId),
	}

	err := global.DB.CreateProductLike(ctx, arg)

	return err
}

func (r *ProductRepository) DeleteProduct(ctx *gin.Context, id int64) error {
	err := global.DB.DeleteProductById(ctx, id)

	return err
}

func (r *ProductRepository) CheckProduct(ctx *gin.Context, id int64) error {
	_, err := global.DB.CheckProduct(ctx, id)

	return err
}

func (r *ProductRepository) DeleteLikeProduct(ctx *gin.Context, id int) error {
	err := global.DB.DeleteLikedProductByUserId(ctx, int32(id))

	return err
}

func (r *ProductRepository) DeleteManyProduct(ctx *gin.Context, arrayId []int64) error {
	err := global.DB.DeleteManyProductsByIds(ctx, arrayId)

	return err
}
