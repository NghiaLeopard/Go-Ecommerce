package response

import (
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
)

type IAddressesResponse struct {
	Address     string `json:"address"`
	City        string `json:"city"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MiddleName  string `json:"middleName"`
	IsDefault   bool   `json:"isDefault"`
}

type IRoleResponse struct {
	Id         int      `json:"_id"`
	Name       string   `json:"name"`
	Permission []string `json:"permissions"`
}

type UserResponse struct {
	Id                   int                  `json:"_id"`
	Email                string               `json:"email"`
	ResetToken           string               `json:"resetToken"`
	Status               db.UsersStatus       `json:"status"`
	Address              string               `json:"address"`
	Avatar               string               `json:"avatar"`
	PhoneNumber          int                  `json:"phoneNumber"`
	Role                 IRoleResponse        `json:"role"`
	FirstName            string               `json:"firstName"`
	LastName             string               `json:"lastName"`
	MiddleName           string               `json:"middleName"`
	City                 int                  `json:"city"`
	LikeProducts         int                  `json:"likeProducts"`
	ViewedProducts       int                  `json:"viewedProducts"`
	Addresses            []IAddressesResponse `json:"addresses"`
	ResetTokenExpiration time.Time            `json:"resetTokenExpiration"`
	Create_at            time.Time            `json:"create_at"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         UserResponse
}
