package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	AuthRepo IRepository.IAuth
}

func NewAuthUseCase(authRepo IRepository.IAuth) IUseCase.IAuthUseCase {
	return &AuthUseCase{AuthRepo: authRepo}
}

// LoginUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LoginUseCase(ctx *gin.Context, email string, password string) (response.ILoginResponse, error, int) {
	user, err := a.AuthRepo.GetUserByEmail(ctx, email)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ILoginResponse{}, fmt.Errorf("account is not exist: %w", err), 400
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ILoginResponse{}, fmt.Errorf("password is not correct: %w", err), 400
	}

	accessToken, _, err := global.Token.CreateTokenPaseto(int(user.ID), user.Permission, global.Config.Access_token)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ILoginResponse{}, fmt.Errorf("generate access token false: %w", err), 500
	}

	refreshToken, _, err := global.Token.CreateTokenPaseto(int(user.ID), user.Permission, global.Config.Refresh_token)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.ILoginResponse{}, fmt.Errorf("generate refresh token false: %w", err), 500
	}

	data := response.ILoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: response.IUserResponse{
			Id:          int(user.ID),
			Email:       user.Email,
			ResetToken:  "",
			Address:     user.Address.String,
			Status:      user.Status.UsersStatus,
			Avatar:      user.Avatar.String,
			PhoneNumber: int(user.PhoneNumber.Int64),
			Role: response.IRoleResponse{
				Id:         int(user.ID_2),
				Name:       user.Name,
				Permission: user.Permission,
			},
			FirstName:            user.FirstName.String,
			LastName:             user.LastName.String,
			MiddleName:           user.MiddleName.String,
			City:                 int(user.City.Int64),
			LikeProducts:         user.LikeProducts,
			ViewedProducts:       user.ViewedProducts,
			ResetTokenExpiration: user.ResetTokenExpiration.Time,
			Create_at:            user.CreateAt,
		},
	}

	global.Logger.Info("Login", zap.String("Status", "success"))
	return data, nil, 0
}

// RegisterUseCase implements IUseCase.IAuthUseCase.
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

// ChangePasswordUseCase implements IUseCase.IAuthUseCase.
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

// LogoutUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LogoutUseCase(ctx *gin.Context) (error, int) {
	// TODO: blackList
	return nil, 200
}

// ForgotPasswordUseCase implements IUseCase.IAuthUseCase.
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

	arg := db.SaveResetTokenParams{
		ID: user.ID,
		// ResetToken: token,
	}

	global.DB.SaveResetToken(ctx, arg)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("send email fail"), 400
	}

	global.Logger.Info("Forgot password", zap.String("Status", "success"))
	return nil, 200
}

// ResetPasswordUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) ResetPasswordUseCase(ctx *gin.Context, newPassword string, secretKey string) (error, int) {
	payload, err := global.Token.VerifyTokenPaseto(secretKey)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("token is invalid"), 400
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

// GetAuthMeUserCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) GetAuthMeUserCase(ctx *gin.Context) (response.IAuthMe, error, int) {
	payload := ctx.MustGet(middleware.AuthorizationKey).(*token.Payload)

	user, err := a.AuthRepo.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return response.IAuthMe{}, fmt.Errorf("get auth me fail"), 500
	}

	data := response.IAuthMe{
		Id:      int(user.ID),
		Email:   user.Email,
		Address: user.Address.String,
		Status:  user.Status.UsersStatus,
		Role: response.IRoleResponse{
			Id:         int(user.ID_2),
			Name:       user.Name,
			Permission: user.Permission,
		},
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		MiddleName:  user.MiddleName.String,
		City:        int(user.City.Int64),
		PhoneNumber: int(user.PhoneNumber.Int64),
		Avatar:      user.Avatar.String,
		Addresses:   []response.IAddressesResponse{},
		Create_at:   user.CreateAt,
	}

	global.Logger.Info("get auth me", zap.String("Status", "success"))
	return data, nil, 200
}
