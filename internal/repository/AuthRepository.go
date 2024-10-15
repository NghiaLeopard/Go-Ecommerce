package repository

import (
	"database/sql"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type AuthRepository struct {
}

func (a *AuthRepository) FindUserById(ctx *gin.Context, id int64) error {
	err := global.DB.FindUserById(ctx, id)

	return err
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
func (a *AuthRepository) UpdateAuthMe(ctx *gin.Context, req IRequest.UpdateAuthMe, id int64) (user db.User, err error) {

	arg := db.UpdateAuthMeParams{
		Avatar:      sql.NullString{String: req.Avatar, Valid: true},
		Address:     sql.NullString{String: req.Address, Valid: true},
		FirstName:   sql.NullString{String: req.FirstName, Valid: true},
		LastName:    sql.NullString{String: req.LastName, Valid: true},
		MiddleName:  sql.NullString{String: req.MiddleName, Valid: true},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: true},
		City:        sql.NullInt64{Int64: req.City, Valid: req.City != 0},
		ID:          id,
		Image:       sql.NullString{String: req.Image, Valid: true},
	}
	user, err = global.DB.UpdateAuthMe(ctx, arg)

	return
}
