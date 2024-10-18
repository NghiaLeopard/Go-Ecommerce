package repository

import (
	"database/sql"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

func NewUserRepository() IRepository.User {
	return &UserRepository{}
}

func (c *UserRepository) CreateUserAdmin(ctx *gin.Context, req IRequest.CreateUser) (User db.User, err error) {
	arg := db.CreateUserAdminParams{
		Email:       req.Email,
		Password:    req.Password,
		FirstName:   sql.NullString{String: req.FirstName, Valid: true},
		LastName:    sql.NullString{String: req.LastName, Valid: true},
		MiddleName:  sql.NullString{String: req.MiddleName, Valid: true},
		Avatar:      sql.NullString{String: req.Avatar, Valid: true},
		Address:     sql.NullString{String: req.Address, Valid: true},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: true},
		Role:        sql.NullInt64{Int64: req.Role, Valid: req.Role != 0},
		City:        sql.NullInt64{Int64: req.City, Valid: req.City != 0},
	}

	User, err = global.DB.CreateUserAdmin(ctx, arg)

	return
}

func (c *UserRepository) UpdateUser(ctx *gin.Context, id int64, body IRequest.GetBodyUpdateUser) (User db.User, err error) {
	arg := db.UpdateUserAdminParams{
		FirstName:   sql.NullString{String: body.FirstName, Valid: body.FirstName != ""},
		LastName:    sql.NullString{String: body.LastName, Valid: body.LastName != ""},
		MiddleName:  sql.NullString{String: body.MiddleName, Valid: body.MiddleName != ""},
		Avatar:      sql.NullString{String: body.Avatar, Valid: body.Avatar != ""},
		Address:     sql.NullString{String: body.Address, Valid: body.Address != ""},
		PhoneNumber: sql.NullString{String: body.PhoneNumber, Valid: body.PhoneNumber != ""},
		Role:        sql.NullInt64{Int64: body.Role, Valid: body.Role != 0},
		City:        sql.NullInt64{Int64: body.City, Valid: body.City != 0},
		Status:      db.UsersStatus(body.Status),
		ID:          id,
	}

	User, err = global.DB.UpdateUserAdmin(ctx, arg)

	return
}

func (c *UserRepository) GetUserById(ctx *gin.Context, id int64) (User db.GetUserAdminByIdRow, err error) {
	User, err = global.DB.GetUserAdminById(ctx, id)

	return
}

func (c *UserRepository) GetAllUser(ctx *gin.Context, page int32, limit int32, search string, order string) (User []db.ListUserAdminRow, err error) {

	offset := limit * (page - 1)

	arg := db.ListUserAdminParams{
		LimitOpt:  limit,
		OffsetOpt: offset,
		OrderBy:   order,
		Search:    search,
	}
	User, err = global.DB.ListUserAdmin(ctx, arg)

	return
}

func (c *UserRepository) GetUserByEmail(ctx *gin.Context, email string) (User db.GetUserByEmailRow, err error) {
	User, err = global.DB.GetUserByEmail(ctx, email)

	return
}

func (c *UserRepository) DeleteUserById(ctx *gin.Context, id int64) (err error) {
	err = global.DB.DeleteUserAdminById(ctx, id)

	return
}

func (c *UserRepository) DeleteManyUserByIds(ctx *gin.Context, arrayID []int64) (err error) {

	err = global.DB.DeleteManyUserAdminByIds(ctx, arrayID)

	return
}
