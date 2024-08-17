package usecase

import (
	"database/sql"
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthUseCase struct {
	DB     *db.Queries
	Config config.Config
	Token  token.Maker
}

// LogoutUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LogoutUseCase(ctx *gin.Context) (error, int) {
	// TODO: blackList
	return nil, 200
}

func NewAuthUseCase(sqlDB *sql.DB, config config.Config, token token.Maker) IUseCase.IAuthUseCase {
	return &AuthUseCase{DB: db.New(sqlDB), Config: config, Token: token}
}

// LoginUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LoginUseCase(ctx *gin.Context, email string, password string) (response.LoginResponse, error, int) {
	user, err := a.DB.FindEmail(ctx, email)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("account is not exist: %w", err), 400
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("password is not correct: %w", err), 400
	}

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("New paseto token false: %w", err), 500
	}

	accessToken, _, err := a.Token.CreateTokenPaseto(int(user.ID), a.Config.Access_token)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("Generate access token false: %w", err), 500
	}

	refreshToken, _, err := a.Token.CreateTokenPaseto(int(user.ID), a.Config.Refresh_token)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("Generate refresh token false: %w", err), 500
	}

	data := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: response.UserResponse{
			Id:                   int(user.ID),
			Email:                user.Email,
			ResetToken:           "",
			Address:              user.Address.String,
			Status:               user.Status.UsersStatus,
			Avatar:               user.Avatar.String,
			PhoneNumber:          int(user.PhoneNumber.Int64),
			Role:                 int(user.Role.Int64),
			FirstName:            user.FirstName.String,
			LastName:             user.LastName.String,
			MiddleName:           user.MiddleName.String,
			City:                 int(user.City.Int64),
			LikeProducts:         int(user.LikeProducts.Int64),
			ViewedProducts:       int(user.ViewedProducts.Int64),
			ResetTokenExpiration: user.ResetTokenExpiration.Time,
			Create_at:            user.CreateAt,
		},
	}

	return data, nil, 0
}

// RegisterUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) RegisterUseCase(ctx *gin.Context, email string, password string) error {
	hashPassword, err := utils.HashPassword(password)

	if err != nil {
		return fmt.Errorf("hash password fail: %w", err)
	}

	arg := db.CreateUserParams{
		Email:    email,
		Password: hashPassword,
	}

	_, err = a.DB.CreateUser(ctx, arg)

	if err != nil {
		return fmt.Errorf("create user fail: %w", err)
	}

	return nil
}
