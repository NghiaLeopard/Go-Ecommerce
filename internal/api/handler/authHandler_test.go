package handler

import (
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/db/mock"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/golang/mock/gomock"
)

func TestLoginUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockQuerier := mock.NewMockQuerier(ctrl)

	arg := db.CreateUserParams{
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	mockQuerier.EXPECT().CreateUser(gomock.Any(), gomock.Eq(arg)).Times(1).Return()

}
