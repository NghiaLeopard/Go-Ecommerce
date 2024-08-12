package usecase

import (
	"database/sql"
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthUseCase struct {
	DB *db.Queries
}

func NewAuthUseCase(sqlDB *sql.DB) IUseCase.IAuthUseCase {
	return &AuthUseCase{DB: db.New(sqlDB)}
}

// LoginUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LoginUseCase(ctx *gin.Context,email string, password string) (response.LoginResponse, error) {
	user,err := a.DB.FindEmail(ctx,email)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("account is not exist: %w", err)
	}

	err = utils.CheckPassword(user.Password,password)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("password is not correct: %w", err)
	}

	data := response.LoginResponse{}

	return data,nil
}

// RegisterUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) RegisterUseCase(ctx *gin.Context,email string, password string) error {
	hashPassword,err := utils.HashPassword(password)

	if err != nil {
		return fmt.Errorf("hash password fail: %w",err)
	}

	arg := db.CreateUserParams{
		Email: email,
		Password: hashPassword,
	}

	_,err =  a.DB.CreateUser(ctx,arg)

	if err != nil {
		return fmt.Errorf("create user fail: %w",err)
	}

	return nil
}


