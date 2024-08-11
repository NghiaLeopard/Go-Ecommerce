package IUseCase

import "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"

type IAuthUseCase interface {
	LoginUseCase(email string, password string) (response.AuthResponse,error)
	RegisterUseCase(email string, password string) error
}