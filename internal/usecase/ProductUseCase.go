package usecase

import (
	"fmt"
	"sync"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
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
			Discount:          Product.Discount,
			DiscountStartDate: Product.DiscountStartDate,
			DiscountEndDate:   Product.DiscountEndDate,
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

	res := IResponse.Product{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount,
		DiscountStartDate: Product.DiscountStartDate,
		DiscountEndDate:   Product.DiscountEndDate,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        0,
		Views:             0,
		LikedBy:           []int32{},
		UniqueViews:       []int32{},
		CreateAt:          Product.CreateAt,
	}

	return res, nil, 201

}

func (c *ProductUseCase) GetAllProductAdminUseCase(ctx *gin.Context, req IRequest.GetAllProductAdmin) (IResponse.GetAllProductAdmin, error, int) {
	product, err := c.ProductRepo.GetAllProductAdmin(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllProductAdmin{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0}, fmt.Errorf("get Product is not exist"), 400
	}

	if len(product) == 0 {
		return IResponse.GetAllProductAdmin{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), product[0].TotalCount)

	return IResponse.GetAllProductAdmin{
		Products:   product,
		TotalCount: product[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

func (c *ProductUseCase) GetAllProductPublicUseCase(ctx *gin.Context, req IRequest.GetAllProductPublic) (IResponse.GetAllProductPublic, error, int) {
	product, err := c.ProductRepo.GetAllProductPublic(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllProductPublic{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0,
		}, fmt.Errorf("get Product is not exist"), 400
	}

	if len(product) == 0 {
		return IResponse.GetAllProductPublic{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), product[0].TotalCount)

	return IResponse.GetAllProductPublic{
		Products:   product,
		TotalCount: product[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

func (c *ProductUseCase) GetAllProductMeLikedUseCase(ctx *gin.Context, req IRequest.GetAllProductLiked) (IResponse.GetAllMeLiked, error, int) {
	payload := ctx.MustGet(constant.AuthorizationKey).(*token.Payload)

	product, err := c.ProductRepo.GetAllProductMeLiked(ctx, req, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllMeLiked{}, fmt.Errorf("get Product is not exist"), 401
	}
	if len(product) == 0 {
		return IResponse.GetAllMeLiked{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), product[0].TotalCount)

	return IResponse.GetAllMeLiked{
		Products:   product,
		TotalCount: product[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

func (c *ProductUseCase) GetAllProductMeViewedUseCase(ctx *gin.Context, req IRequest.GetAllProductViewed) (IResponse.GetAllMeViewed, error, int) {
	payload := ctx.MustGet(constant.AuthorizationKey).(*token.Payload)

	product, err := c.ProductRepo.GetAllProductMeViewed(ctx, req, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllMeViewed{}, fmt.Errorf("get Product is not exist"), 401
	}

	if len(product) == 0 {
		return IResponse.GetAllMeViewed{
			Products:   product,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), product[0].TotalCount)

	return IResponse.GetAllMeViewed{
		Products:   product,
		TotalCount: product[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

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
		Image:             Product.Image,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount,
		DiscountStartDate: Product.DiscountStartDate,
		DiscountEndDate:   Product.DiscountEndDate,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        Product.TotalLikes,
		Views:             Product.Views,
		LikedBy:           Product.LikedBy,
		UniqueViews:       Product.UniqueViews,
		CreateAt:          Product.CreateAt,
	}

	return res, nil, 200
}

func (c *ProductUseCase) GetProductRelatedUseCase(ctx *gin.Context, req IRequest.GetAllProductRelated) (IResponse.GetAllProductRelated, error, int) {
	product, err := c.ProductRepo.GetProductTypeBySlug(ctx, req.Slug)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllProductRelated{}, fmt.Errorf("product is not exist"), 401
	}

	productRelate, err := c.ProductRepo.GetAllProductRelated(ctx, req, product.ID, product.Type)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllProductRelated{}, fmt.Errorf("server error"), 500
	}

	if len(productRelate) == 0 {
		return IResponse.GetAllProductRelated{
			Products:   productRelate,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), productRelate[0].TotalCount)

	return IResponse.GetAllProductRelated{
		Products:   productRelate,
		TotalCount: productRelate[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200

}

func (c *ProductUseCase) GetProductBySlugUseCase(ctx *gin.Context, slug string, isViewed bool) (IResponse.GetProduct, error, int) {
	Product, err := c.ProductRepo.GetProductBySlug(ctx, slug)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetProduct{}, fmt.Errorf("get Product is not exist"), 401
	}

	if isViewed {
		var wg sync.WaitGroup
		wg.Add(1)
		view := Product.Views + 1
		go c.ProductRepo.UpdateViewProduct(ctx, Product.ID, view, &wg)

		if checkValue, exists := ctx.Get(constant.AuthorizationKey); exists {
			value, ok := checkValue.(*token.Payload)
			if !ok {
				global.Logger.Error("Failed to cast value to *token.Payload")
				return IResponse.GetProduct{}, fmt.Errorf("invalid authorization token"), 401
			}

			boolRedis, err := c.RedisProduct.CheckProductUniqueView(ctx, Product.ID, value.Id)
			if !boolRedis && err == nil {
				wg.Add(1)
				go c.ProductRepo.UpdateUniqueView(ctx, Product.ID, value.Id, &wg)
				err := c.RedisProduct.SetProductUniqueView(ctx, Product.ID, value.Id)

				if err != nil {
					global.Logger.Error(err.Error(), zap.String("Status", "Error"))
					return IResponse.GetProduct{}, fmt.Errorf("redis set unique product"), 500
				}

			}
		}

		wg.Wait()
	}

	res := IResponse.GetProduct{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount,
		DiscountStartDate: Product.DiscountStartDate,
		DiscountEndDate:   Product.DiscountEndDate,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        Product.TotalLikes,
		Views:             Product.Views,
		LikedBy:           Product.LikedBy,
		UniqueViews:       Product.UniqueViews,
		CreateAt:          Product.CreateAt,
		Image:             Product.Image,
	}

	return res, nil, 200
}

func (c *ProductUseCase) GetProductPublicByIdUseCase(ctx *gin.Context, id int64, isViewed bool) (IResponse.GetProduct, error, int) {
	Product, err := c.ProductRepo.GetProductPublicById(ctx, id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetProduct{}, fmt.Errorf("get Product is not exist"), 401
	}

	if isViewed {
		var wg sync.WaitGroup
		wg.Add(1)
		view := Product.Views + 1
		go c.ProductRepo.UpdateViewProduct(ctx, Product.ID, view, &wg)

		if checkValue, exists := ctx.Get(constant.AuthorizationKey); exists {
			value, ok := checkValue.(*token.Payload)
			if !ok {
				global.Logger.Error("Failed to cast value to *token.Payload")
				return IResponse.GetProduct{}, fmt.Errorf("invalid authorization token"), 401
			}

			boolRedis, err := c.RedisProduct.CheckProductUniqueView(ctx, Product.ID, value.Id)
			if !boolRedis && err == nil {
				wg.Add(1)
				go c.ProductRepo.UpdateUniqueView(ctx, Product.ID, value.Id, &wg)
				err := c.RedisProduct.SetProductUniqueView(ctx, Product.ID, value.Id)

				if err != nil {
					global.Logger.Error(err.Error(), zap.String("Status", "Error"))
					return IResponse.GetProduct{}, fmt.Errorf("redis set unique product"), 500
				}

			}
		}

		wg.Wait()
	}

	res := IResponse.GetProduct{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount,
		DiscountStartDate: Product.DiscountStartDate,
		DiscountEndDate:   Product.DiscountEndDate,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		TotalLikes:        Product.TotalLikes,
		Views:             Product.Views,
		LikedBy:           Product.LikedBy,
		UniqueViews:       Product.UniqueViews,
		CreateAt:          Product.CreateAt,
	}

	return res, nil, 200
}

func (c *ProductUseCase) UpdateProductUseCase(ctx *gin.Context, id int64, body IRequest.UpdateProduct) (IResponse.UpdateProduct, error, int) {

	_, err := global.DB.CheckProduct(ctx, id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.UpdateProduct{}, fmt.Errorf("product is not exist"), 409
	}

	Product, err := c.ProductRepo.UpdateProduct(ctx, id, body)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.UpdateProduct{}, fmt.Errorf("update Product is fail"), 400
	}

	res := IResponse.UpdateProduct{
		Id:                Product.ID,
		Name:              Product.Name,
		Slug:              Product.Slug,
		Price:             Product.Price,
		CountInStock:      Product.CountInStock,
		Description:       Product.Description,
		Discount:          Product.Discount,
		DiscountStartDate: Product.DiscountStartDate,
		DiscountEndDate:   Product.DiscountEndDate,
		Type:              Product.Type,
		Location:          Product.Location,
		Status:            Product.Status,
		Views:             Product.Views,
		CreateAt:          Product.CreateAt,
	}

	return res, nil, 200
}

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

func (c *ProductUseCase) LikeProductUseCase(ctx *gin.Context, productId int64) (error, int) {
	payload := ctx.MustGet(constant.AuthorizationKey).(*token.Payload)

	boolLike, err := c.RedisProduct.CheckLikeProduct(ctx, productId, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("redis error"), 500
	}

	if boolLike {
		global.Logger.Error("userId is exist in productLike", zap.String("Status", "Error"))
		return fmt.Errorf("userId is exist in productLike"), 400
	}

	err = c.RedisProduct.SetLikeProduct(ctx, productId, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("redis error"), 500
	}

	err = c.ProductRepo.UpdateLikeProduct(ctx, productId, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("update product fail"), 400
	}

	return nil, 201
}

func (c *ProductUseCase) UnLikeProductUseCase(ctx *gin.Context, productId int64) (error, int) {
	payload := ctx.MustGet(constant.AuthorizationKey).(*token.Payload)

	boolLike, err := c.RedisProduct.CheckLikeProduct(ctx, productId, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("redis error"), 500
	}

	if !boolLike {
		global.Logger.Error("deleted user liked Product", zap.String("Status", "Error"))
		return fmt.Errorf("deleted user liked Product"), 400
	}
	err = c.RedisProduct.DeleteLikeProduct(ctx, productId, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("redis error"), 500
	}

	err = c.ProductRepo.DeleteLikeProduct(ctx, payload.Id)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Product is not exist"), 400
	}

	return nil, 201
}
