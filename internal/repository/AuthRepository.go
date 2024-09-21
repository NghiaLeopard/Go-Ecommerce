package repository

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type AuthRepository struct {
}

func NewAuthRepository() IRepository.Auth {
	return &AuthRepository{}
}

// CreateUser implements IRepository.IAuth.
func (a *AuthRepository) CreateUser(ctx *gin.Context, arg db.CreateUserParams) (user db.User, err error) {
	user, err = global.DB.CreateUser(ctx, arg)
	return
}

// DeleteUser implements IRepository.IAuth.
func (a *AuthRepository) DeleteUser(ctx *gin.Context, id int64) (err error) {
	err = global.DB.DeleteUser(ctx, id)

	return
}

// GetUserByEmail implements IRepository.IAuth.
func (a *AuthRepository) GetUserByEmail(ctx *gin.Context, email string) (user db.GetUserByEmailRow, err error) {
	user, err = global.DB.GetUserByEmail(ctx, email)

	return
}

// GetUserById implements IRepository.IAuth.
func (a *AuthRepository) GetUserById(ctx *gin.Context, id int64) (user db.GetUserByIdRow, err error) {
	user, err = global.DB.GetUserById(ctx, id)

	return

}

// InitDefaultAdmin implements IRepository.IAuth.
func (a *AuthRepository) InitDefaultAdmin(ctx *gin.Context, arg db.InitDefaultAdminParams) (user db.User, err error) {
	user, err = global.DB.InitDefaultAdmin(ctx, arg)

	return
}

// UpdatePasswordUser implements IRepository.IAuth.
func (a *AuthRepository) UpdatePasswordUser(ctx *gin.Context, arg db.UpdatePasswordUserParams) (err error) {
	err = global.DB.UpdatePasswordUser(ctx, arg)

	return
}

// UpdateUser implements IRepository.IAuth.
func (a *AuthRepository) UpdateUser(ctx *gin.Context, arg db.UpdateUserParams) (user db.User, err error) {
	user, err = global.DB.UpdateUser(ctx, arg)

	return
}
