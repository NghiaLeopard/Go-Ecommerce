package IResponse

import (
	"encoding/json"
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type User struct {
	Id                   int64           `json:"_id"`
	Email                string          `json:"email"`
	ResetToken           string          `json:"resetToken"`
	Status               db.UsersStatus  `json:"status"`
	Address              string          `json:"address"`
	Avatar               string          `json:"avatar"`
	PhoneNumber          int             `json:"phoneNumber"`
	Role                 Role            `json:"role"`
	FirstName            string          `json:"firstName"`
	LastName             string          `json:"lastName"`
	MiddleName           string          `json:"middleName"`
	City                 int             `json:"city"`
	LikeProducts         []int64         `json:"likeProducts"`
	ViewedProducts       []int64         `json:"viewedProducts"`
	Addresses            json.RawMessage `json:"addresses"`
	ResetTokenExpiration time.Time       `json:"resetTokenExpiration"`
	Create_at            time.Time       `json:"create_at"`
}

type AuthMe struct {
	Id          int64           `json:"_id"`
	Email       string          `json:"email"`
	Status      db.UsersStatus  `json:"status"`
	Address     string          `json:"address"`
	Avatar      string          `json:"avatar"`
	PhoneNumber int             `json:"phoneNumber"`
	Role        Role            `json:"role"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	MiddleName  string          `json:"middleName"`
	City        int             `json:"city"`
	Addresses   json.RawMessage `json:"addresses"`
	Create_at   time.Time       `json:"create_at"`
}

type UpdateAuthMe struct {
	Id          int64           `json:"_id" swaggertype:"integer"`
	Email       string          `json:"email"`
	Status      db.UsersStatus  `json:"status"`
	Address     string          `json:"address" swaggertype:"string"`
	Avatar      string          `json:"avatar" swaggertype:"string"`
	PhoneNumber int64           `json:"phoneNumber" swaggertype:"integer"`
	Role        int64           `json:"role" `
	FirstName   string          `json:"firstName" swaggertype:"string"`
	LastName    string          `json:"lastName" swaggertype:"string"`
	MiddleName  string          `json:"middleName" swaggertype:"string"`
	City        int64           `json:"city" swaggertype:"integer"`
	Addresses   json.RawMessage `json:"addresses"`
	Create_at   time.Time       `json:"create_at"`
}

type Login struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetAccessToken struct {
	AccessToken string `json:"access_token"`
}
