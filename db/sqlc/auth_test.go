package db

import (
	"context"
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreateAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user := CreateRandomUser(t)
	getUser, err := testQuery.GetUserById(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.Equal(t, user.Email, getUser.Email)
	require.Equal(t, user.Password, getUser.Password)
}

func TestGetUserByEmail(t *testing.T) {
	user := CreateRandomUser(t)
	getUser, err := testQuery.GetUserByEmail(context.Background(), user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.Equal(t, user.Email, getUser.Email)
	require.Equal(t, user.Password, getUser.Password)
}

func TestUpdatePasswordUser(t *testing.T) {
	user := CreateRandomUser(t)

	arg := UpdatePasswordUserParams{
		Password: user.Password,
		ID:       user.ID,
	}

	err := testQuery.UpdatePasswordUser(context.Background(), arg)

	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	err := testQuery.DeleteUser(context.Background(), 3)

	require.NoError(t, err)
}
