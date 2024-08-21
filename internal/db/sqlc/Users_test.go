package db

import (
	"context"
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreateAt)
}

func TestDeleteUser(t *testing.T) {

	err := testQuery.DeleteUser(context.Background(), 2)

	require.NoError(t, err)
}
