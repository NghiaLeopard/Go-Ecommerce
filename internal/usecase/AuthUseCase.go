package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/gmail"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthUseCase struct {
	DB          *db.Queries
	Config      config.Config
	Token       token.Maker
	EmailSender gmail.Sender
}

func NewAuthUseCase(db *db.Queries, config config.Config, token token.Maker, email gmail.Sender) IUseCase.IAuthUseCase {
	return &AuthUseCase{DB: db, Config: config, Token: token, EmailSender: email}
}

// LoginUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LoginUseCase(ctx *gin.Context, email string, password string) (response.LoginResponse, error, int) {
	user, err := a.DB.GetUserByEmail(ctx, email)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("account is not exist: %w", err), 400
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("password is not correct: %w", err), 400
	}

	accessToken, _, err := a.Token.CreateTokenPaseto(int(user.ID), user.Permission, a.Config.Access_token)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("generate access token false: %w", err), 500
	}

	refreshToken, _, err := a.Token.CreateTokenPaseto(int(user.ID), user.Permission, a.Config.Refresh_token)

	if err != nil {
		return response.LoginResponse{}, fmt.Errorf("generate refresh token false: %w", err), 500
	}

	data := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: response.UserResponse{
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

// ChangePasswordUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) ChangePasswordUseCase(ctx *gin.Context, currentPassword string, newPassword string) (error, int) {
	payload := ctx.MustGet(middleware.AuthorizationKey).(*token.Payload)

	user, err := a.DB.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		return fmt.Errorf("get user db fail"), 500
	}

	err = utils.CheckPassword(user.Password, currentPassword)

	if err != nil {
		return fmt.Errorf("password is not correct"), 400
	}

	hashNewPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		return fmt.Errorf("hash password fail"), 500
	}

	arg := db.UpdatePasswordUserParams{
		ID:       user.ID,
		Password: hashNewPassword,
	}

	err = a.DB.UpdatePasswordUser(ctx, arg)

	if err != nil {
		return fmt.Errorf("update user fail"), 500
	}

	return nil, 200
}

// LogoutUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LogoutUseCase(ctx *gin.Context) (error, int) {
	// TODO: blackList
	return nil, 200
}

// ForgotPasswordUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) ForgotPasswordUseCase(ctx *gin.Context, email string) (error, int) {
	user, err := a.DB.GetUserByEmail(ctx, email)

	if err != nil {
		return fmt.Errorf("email is not exist"), 400
	}

	token, _, err := a.Token.CreateTokenPaseto(int(user.ID), []string{}, a.Config.ForgotPasswordToken)

	if err != nil {
		return fmt.Errorf("email is not exist"), 400
	}

	textEmail := fmt.Sprintf("%s?%s", a.Config.AppUrlFE, token)
	var subject = "Send link forgot password"

	err = a.EmailSender.SenderEmail([]string{email}, subject, []byte(textEmail), nil, nil)

	if err != nil {
		return fmt.Errorf("send email fail"), 400
	}

	return nil, 200
}

// ResetPasswordUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) ResetPasswordUseCase(ctx *gin.Context, newPassword string, secretKey string) (error, int) {
	payload, err := a.Token.VerifyTokenPaseto(secretKey)

	if err != nil {
		return fmt.Errorf("token is invalid"), 400
	}

	user, err := a.DB.GetUserById(ctx, int64(payload.Id))

	if err != nil {
		return fmt.Errorf("query sql"), 500
	}

	hashPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		return fmt.Errorf("err hash password"), 500
	}

	arg := db.UpdatePasswordUserParams{
		Password: hashPassword,
		ID:       user.ID,
	}

	err = a.DB.UpdatePasswordUser(ctx, arg)

	if err != nil {
		return fmt.Errorf("update err"), 500
	}

	return nil, 200
}
