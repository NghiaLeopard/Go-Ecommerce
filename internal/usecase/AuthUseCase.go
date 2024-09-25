package usecase

import (
	"fmt"
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	AuthorizationHeader = "authorization"
	AuthorizationType   = "Bearer"
)

type AuthUseCase struct {
	AuthRepo       IRepository.Auth
	RedisTokenRepo IRepository.RedisToken
}

func NewAuthUseCase(authRepo IRepository.Auth, redisToken IRepository.RedisToken) IUseCase.Auth {
	return &AuthUseCase{AuthRepo: authRepo, RedisTokenRepo: redisToken}
}

func (a *AuthUseCase) RefreshTokenUseCase(ctx *gin.Context) (IResponse.GetAccessToken, error, int) {
	authorization := ctx.GetHeader(AuthorizationHeader)

	if len(authorization) == 0 {
		global.Logger.Error("please provide authorization", zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf("please provide authorization"), 401
	}

	fields := strings.Fields(authorization)

	if len(fields) < 2 {
		global.Logger.Error("invalid format header", zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf("invalid format header"), 401

	}

	if fields[0] != AuthorizationType {
		global.Logger.Error("invalid type header", zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf("invalid type header"), 401

	}

	payload, err := global.Token.VerifyTokenPaseto(fields[1])

	if err != nil {
		global.Logger.Error("Verify token invalid", zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf(err.Error()), 401
	}

	err = a.RedisTokenRepo.CheckRefreshToken(ctx, int64(payload.Id), fields[1])

	if err != nil {
		global.Logger.Error("Refresh token is invalid", zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf(err.Error()), 401
	}

	accessToken, _, err := global.Token.CreateTokenPaseto(int(payload.Id), payload.Permissions, global.Config.Access_token)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAccessToken{}, fmt.Errorf("generate access token false: %w", err), 500
	}

	return IResponse.GetAccessToken{
		AccessToken: accessToken,
	}, nil, 201
}

func (a *AuthUseCase) LoginUseCase(ctx *gin.Context, email string, password string) (IResponse.Login, error, int) {
	user, err := a.AuthRepo.GetUserByEmail(ctx, email)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Login{}, fmt.Errorf("account is not exist: %w", err), 400
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Login{}, fmt.Errorf("password is not correct: %w", err), 400
	}

	accessToken, _, err := global.Token.CreateTokenPaseto(int(user.ID), user.Permission, global.Config.Access_token)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Login{}, fmt.Errorf("generate access token false: %w", err), 500
	}

	refreshToken, _, err := global.Token.CreateTokenPaseto(int(user.ID), user.Permission, global.Config.Refresh_token)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Login{}, fmt.Errorf("generate refresh token false: %w", err), 500
	}

	err = a.RedisTokenRepo.SetRefreshToken(ctx, user.ID, refreshToken, 0)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Login{}, fmt.Errorf("set redis refresh token false: %w", err), 500
	}

	data := IResponse.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: IResponse.User{
			Id:          user.ID,
			Email:       user.Email,
			Address:     user.Address.String,
			Status:      user.Status.UsersStatus,
			Avatar:      user.Avatar.String,
			PhoneNumber: int(user.PhoneNumber.Int64),
			Role: IResponse.Role{
				Id:         user.ID_2,
				Name:       user.Name,
				Permission: user.Permission,
			},
			FirstName:      user.FirstName.String,
			LastName:       user.LastName.String,
			MiddleName:     user.MiddleName.String,
			City:           int(user.City.Int64),
			LikeProducts:   user.LikeProducts,
			ViewedProducts: user.ViewedProducts,
			Create_at:      user.CreateAt,
		},
	}

	global.Logger.Info("Login", zap.String("Status", "success"))
	return data, nil, 0
}

func (a *AuthUseCase) RegisterUseCase(ctx *gin.Context, email string, password string) error {
	hashPassword, err := utils.HashPassword(password)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("hash password fail: %w", err)
	}

	arg := db.CreateUserParams{
		Email:    email,
		Password: hashPassword,
	}

	_, err = a.AuthRepo.CreateUser(ctx, arg)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("create user fail: %w", err)
	}

	global.Logger.Info("Register", zap.String("Status", "success"))
	return nil
}

func (a *AuthUseCase) ChangePasswordUseCase(ctx *gin.Context, currentPassword string, newPassword string) (error, int) {
	payload := ctx.MustGet(middleware.AuthorizationKey).(*token.Payload)

	user, err := a.AuthRepo.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get user db fail"), 500
	}

	err = utils.CheckPassword(user.Password, currentPassword)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("password is not correct"), 400
	}

	hashNewPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("hash password fail"), 500
	}

	arg := db.UpdatePasswordUserParams{
		ID:       user.ID,
		Password: hashNewPassword,
	}

	err = a.AuthRepo.UpdatePasswordUser(ctx, arg)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("update user fail"), 500
	}

	global.Logger.Info("Change password", zap.String("Status", "success"))
	return nil, 200
}

func (a *AuthUseCase) LogoutUseCase(ctx *gin.Context) (error, int) {
	authorization := ctx.GetHeader(AuthorizationHeader)

	if len(authorization) == 0 {
		global.Logger.Error("please provide authorization", zap.String("Status", "Error"))
		return fmt.Errorf("please provide authorization"), 401
	}

	fields := strings.Fields(authorization)

	if len(fields) < 2 {
		global.Logger.Error("invalid format header", zap.String("Status", "Error"))
		return fmt.Errorf("invalid format header"), 401

	}

	if fields[0] != AuthorizationType {
		global.Logger.Error("invalid type header", zap.String("Status", "Error"))
		return fmt.Errorf("invalid type header"), 401

	}

	payload, err := global.Token.VerifyTokenPaseto(fields[1])

	if err != nil {
		global.Logger.Error("Verify token invalid", zap.String("Status", "Error"))
		return fmt.Errorf(err.Error()), 401
	}

	err = a.RedisTokenRepo.DeleteRefreshToken(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error("Delete refresh token is false", zap.String("Status", "Error"))
		return fmt.Errorf(err.Error()), 400
	}

	err = a.RedisTokenRepo.BlackListToken(ctx, fields[1])

	if err != nil {
		global.Logger.Error("Set blacklist token is false", zap.String("Status", "Error"))
		return fmt.Errorf(err.Error()), 400
	}

	return nil, 200
}

func (a *AuthUseCase) ForgotPasswordUseCase(ctx *gin.Context, email string) (error, int) {
	user, err := a.AuthRepo.GetUserByEmail(ctx, email)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("email is not exist"), 400
	}

	token, _, err := global.Token.CreateTokenPaseto(int(user.ID), []string{}, global.Config.ForgotPasswordToken)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("email is not exist"), 400
	}

	textEmail := fmt.Sprintf("%s?%s", global.Config.AppUrlFE, token)
	var subject = "Send link forgot password"

	err = global.Gmail.SenderEmail([]string{email}, subject, []byte(textEmail), nil, nil)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("send email fail"), 400
	}

	err = a.RedisTokenRepo.SetResetToken(ctx, user.ID, token, global.Config.ForgotPasswordToken)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("save reset token false"), 400
	}

	global.Logger.Info("Forgot password", zap.String("Status", "success"))
	return nil, 200
}

func (a *AuthUseCase) ResetPasswordUseCase(ctx *gin.Context, newPassword string, secretKey string) (error, int) {
	payload, err := global.Token.VerifyTokenPaseto(secretKey)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("token is invalid"), 400
	}

	err = a.RedisTokenRepo.CheckResetToken(ctx, int64(payload.Id), secretKey)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf(err.Error()), 400
	}

	user, err := a.AuthRepo.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("query sql"), 500
	}

	hashPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("err hash password"), 500
	}

	arg := db.UpdatePasswordUserParams{
		Password: hashPassword,
		ID:       user.ID,
	}

	err = a.AuthRepo.UpdatePasswordUser(ctx, arg)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("update err"), 500
	}

	global.Logger.Info("Reset password", zap.String("Status", "success"))
	return nil, 200
}

func (a *AuthUseCase) GetAuthMeUserCase(ctx *gin.Context) (IResponse.AuthMe, error, int) {
	payload := ctx.MustGet(middleware.AuthorizationKey).(*token.Payload)

	err := a.AuthRepo.FindUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.AuthMe{}, fmt.Errorf("User is not exist"), 400
	}

	user, err := a.AuthRepo.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.AuthMe{}, fmt.Errorf("get auth me fail"), 500
	}

	data := IResponse.AuthMe{
		Id:      user.ID,
		Email:   user.Email,
		Address: user.Address.String,
		Status:  user.Status.UsersStatus,
		Role: IResponse.Role{
			Id:         user.ID_2,
			Name:       user.Name,
			Permission: user.Permission,
		},
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		MiddleName:  user.MiddleName.String,
		City:        int(user.City.Int64),
		PhoneNumber: int(user.PhoneNumber.Int64),
		Avatar:      user.Avatar.String,
		Addresses:   []IResponse.Addresses{},
		Create_at:   user.CreateAt,
	}

	global.Logger.Info("get auth me", zap.String("Status", "success"))
	return data, nil, 200
}

func (a *AuthUseCase) UpdateAuthMeUserCase(ctx *gin.Context, req IRequest.UpdateAuthMe) (IResponse.UpdateAuthMe, error, int) {
	payload := ctx.MustGet(middleware.AuthorizationKey).(*token.Payload)

	err := a.AuthRepo.FindUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.UpdateAuthMe{}, fmt.Errorf("User is not exist"), 400
	}

	user, err := a.AuthRepo.UpdateAuthMe(ctx, req, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.UpdateAuthMe{}, fmt.Errorf("Update auth me fail"), 500
	}

	data := IResponse.UpdateAuthMe{
		Id:          user.ID,
		Email:       user.Email,
		Address:     user.Address.String,
		Status:      user.Status.UsersStatus,
		Role:        user.ID,
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		MiddleName:  user.MiddleName.String,
		City:        user.City.Int64,
		PhoneNumber: user.PhoneNumber.Int64,
		Avatar:      user.Avatar.String,
		Addresses:   []IResponse.Addresses{},
		Create_at:   user.CreateAt,
	}

	global.Logger.Info("get auth me", zap.String("Status", "success"))
	return data, nil, 200
}
