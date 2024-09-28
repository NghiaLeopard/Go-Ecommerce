package usecase

import (
	"fmt"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductUseCase struct {
	ProductRepo  IRepository.Product
	RedisProduct IRepository.RedisProduct
}

func NewProductUseCase(productRepo IRepository.Product, redisProduct IRepository.RedisProduct) IUseCase.Product {
	return &ProductUseCase{
		ProductRepo:  productRepo,
		RedisProduct: redisProduct,
	}
}

func (c *ProductUseCase) CreateProduct(ctx *gin.Context, req IRequest.CreateProduct) (IResponse.Product, error, int) {

	if req.Discount > 1 && (time.Now().After(req.DiscountStart)) && (time.Now().After(req.DiscountEndDate)) {
		Product, err := c.ProductRepo.CreateProductDiscount(ctx, req)

		if err != nil {
			global.Logger.Error(err.Error(), zap.String("Status", "Error"))
			return IResponse.Product{}, err, 401
		}

		global.Logger.Info("Create ", zap.String("Status", "success"))
		return IResponse.Product{
			Id:                Product.ID,
			Name:              Product.Name,
			Slug:              Product.Slug,
			Price:             Product.Price,
			CountInStock:      Product.CountInStock,
			Description:       Product.Description,
			Discount:          Product.Discount.Int32,
			DiscountStartDate: Product.DiscountStartDate.Time,
			DiscountEndDate:   Product.DiscountEndDate.Time,
			Type:              Product.Type,
			Location:          Product.Location,
			Status:            Product.Status,
			TotalLikes:        0,
			Views:             0,
			LikedBy:           []int32{},
			UniqueViews:       []int32{},
			CreateAt:          Product.CreateAt,
		}, nil, 201
	}

	Product, err := c.ProductRepo.CreateProductNotDiscount(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Product{}, err, 401
	}

	err = c.RedisProduct.SetViewProduct(ctx, Product.ID)

	if err != nil {
		global.Logger.Error("Redis product error", zap.String("Status", "Error"))
		return IResponse.Product{}, fmt.Errorf("redis product error"), 500
	}

	res := IResponse.Product{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount.Int32,
		DiscountStartDate: Product.DiscountStartDate.Time,
		DiscountEndDate:   Product.DiscountEndDate.Time,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        0,
		Views:             0,
		LikedBy:           []int32{},
		UniqueViews:       []int32{},
		CreateAt:          Product.CreateAt,
	}

	global.Logger.Info("Create ", zap.String("Status", "success"))
	return res, nil, 201

}

// func (c *ProductUseCase) GetAllProductUseCase(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error, int) {
// 	Product, err := c.ProductRepo.GetAllProduct(ctx, req)

// 	if err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		return []db.Product{}, fmt.Errorf("get Product is not exist"), 401
// 	}

// 	return Product, nil, 200
// }

func (c *ProductUseCase) GetProductUseCase(ctx *gin.Context, id int64) (IResponse.GetProduct, error, int) {
	Product, err := c.ProductRepo.GetProductById(ctx, id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetProduct{}, fmt.Errorf("get Product is not exist"), 401
	}

	res := IResponse.GetProduct{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount.Int32,
		DiscountStartDate: Product.DiscountStartDate.Time,
		DiscountEndDate:   Product.DiscountEndDate.Time,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        Product.TotalLikes,
		Views:             Product.Views.Int32,
		LikedBy:           Product.LikedBy,
		UniqueViews:       Product.UniqueViews,
		CreateAt:          Product.CreateAt,
	}

	global.Logger.Info("Create ", zap.String("Status", "success"))
	return res, nil, 201
}

func (c *ProductUseCase) GetProductBySlugUseCase(ctx *gin.Context, slug string, isViewed bool) (IResponse.GetProduct, error, int) {
	Product, err := c.ProductRepo.GetProductBySlug(ctx, slug, isViewed)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetProduct{}, fmt.Errorf("get Product is not exist"), 401
	}

	res := IResponse.GetProduct{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount.Int32,
		DiscountStartDate: Product.DiscountStartDate.Time,
		DiscountEndDate:   Product.DiscountEndDate.Time,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        Product.TotalLikes,
		Views:             Product.Views.Int32,
		LikedBy:           Product.LikedBy,
		UniqueViews:       Product.UniqueViews,
		CreateAt:          Product.CreateAt,
	}

	global.Logger.Info("Create ", zap.String("Status", "success"))
	return res, nil, 201
}

// func (c *ProductUseCase) UpdateProductUseCase(ctx *gin.Context, id int, name string, slug string) (IResponse.Product, error, int) {
// 	idInt64 := int64(id)

// 	Product, err := global.DB.GetProductById(ctx, idInt64)

// 	if err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		return IResponse.Product{}, fmt.Errorf("Product is not exist"), 401
// 	}

// 	Product, err = c.ProductRepo.UpdateProduct(ctx, idInt64, name, slug)

// 	if err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		return IResponse.Product{}, fmt.Errorf("update Product is fail"), 401
// 	}

// 	res := IResponse.Product{
// 		Id:       Product.ID,
// 		Name:     Product.Name,
// 		Slug:     Product.Slug,
// 		CreateAt: Product.CreateAt,
// 		UpdateAt: Product.UpdateAt.Time,
// 	}

// 	return res, nil, 200
// }

func (c *ProductUseCase) DeleteProductUseCase(ctx *gin.Context, id int64) (error, int) {
	err := global.DB.DeleteProductById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Product is not exist"), 401
	}

	return nil, 200
}

func (c *ProductUseCase) DeleteManyProductUseCase(ctx *gin.Context, arrayId []int64) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayId64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayId64[i] = int64(v)
	}

	err := global.DB.DeleteManyProductsByIds(ctx, arrayId64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Product is not exist"), 401
	}

	return nil, 200
}
