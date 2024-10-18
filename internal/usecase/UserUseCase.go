package usecase

import (
	"fmt"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserUseCase struct {
	UserRepo IRepository.User
}

func NewUserUseCase(UserRepo IRepository.User) IUseCase.User {
	return &UserUseCase{UserRepo: UserRepo}
}

func (c *UserUseCase) CreateUserUseCase(ctx *gin.Context, req IRequest.CreateUser) (error, int) {
	_, err := c.UserRepo.GetUserByEmail(ctx, req.Email)

	if err == nil {
		global.Logger.Error("User is  exist", zap.String("Status", "Error"))
		return fmt.Errorf("User is  exist"), 409
	}

	_, err = c.UserRepo.CreateUserAdmin(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return err, 401
	}

	return nil, 201
}

func (c *UserUseCase) GetUserUseCase(ctx *gin.Context, id int) (db.GetUserAdminByIdRow, error, int) {

	User, err := c.UserRepo.GetUserById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return db.GetUserAdminByIdRow{}, fmt.Errorf("get User is not exist"), 401
	}

	return User, nil, 200
}

func (c *UserUseCase) GetAllUserUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllUser, error, int) {
	User, err := c.UserRepo.GetAllUser(ctx, page, limit, search, order)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllUser{}, fmt.Errorf("get User is not exist"), 401
	}

	fmt.Println(User, limit, order)

	if len(User) == 0 {
		return IResponse.GetAllUser{
			Users:      User,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(limit), User[0].TotalCount)

	return IResponse.GetAllUser{
		Users:      User,
		TotalCount: User[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

func (c *UserUseCase) UpdateUserUseCase(ctx *gin.Context, id int, name string) (error, int) {
	idInt64 := int64(id)

	_, err := global.DB.GetUserById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("User is not exist"), 401
	}

	_, err = c.UserRepo.UpdateUser(ctx, idInt64, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("update User is fail"), 401
	}

	return nil, 200
}

func (c *UserUseCase) DeleteUserUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeleteUserAdminById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get User is not exist"), 401
	}

	return nil, 200
}

func (c *UserUseCase) DeleteManyUserUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayInt64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayInt64[i] = int64(v)
	}

	err := global.DB.DeleteManyUserAdminByIds(ctx, arrayInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get User is not exist"), 401
	}

	return nil, 200
}
