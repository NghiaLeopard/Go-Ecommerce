package response

import "time"

type IAddresses struct {
	Address  string `json:"address"`
    City  string `json:"city"`
    PhoneNumber  string `json:"phoneNumber"`
    FirstName  string `json:"firstName"`
    LastName  string `json:"lastName"`
    MiddleName  string `json:"middleName"`
    IsDefault  bool `json:"isDefault"`
}

type UserResponse struct {
	Id                   int `json:"_id"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ResetToken           string `json:"resetToken"`
	Status               string `json:"status"`
	Address              string `json:"address"`
	Avatar               string `json:"avatar"`
	PhoneNumber          int `json:"phoneNumber"`
	Role                 int `json:"role"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	MiddleName           string `json:"middleName"`
	City                 int `json:"city"`
	LikeProducts         int `json:"likeProducts"`
	ViewedProducts       int `json:"viewedProducts"`
	Addresses            []IAddresses `json:"addresses"`
	ResetTokenExpiration time.Time `json:"resetTokenExpiration"`
	Create_at            time.Time `json:"create_at"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	user UserResponse
}