package usecase

import (
	"database/sql"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
)

type AuthUseCase struct {
	DB *db.Queries
}

func NewAuthUseCase(sqlDB *sql.DB) IUseCase.IAuthUseCase {
	return &AuthUseCase{DB: db.New(sqlDB)}
}

// LoginUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) LoginUseCase(email string, password string) (response.AuthResponse, error) {
	panic("unimplemented")
}

// RegisterUseCase implements IUseCase.IAuthUseCase.
func (a *AuthUseCase) RegisterUseCase(email string, password string) error {
	panic("unimplemented")
}


