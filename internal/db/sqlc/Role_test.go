package db

import (
	"context"
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func CreateRoleUser(t *testing.T) Role {
	arg := CreateRoleParams{
		Name:       utils.RandomString(10),
		Permission: []string{},
	}

	role, err := testQuery.CreateRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, role.ID)
	require.Equal(t, arg.Name, role.Name)
	require.Equal(t, arg.Permission, role.Permission)

	return role
}

func TestCreateRole(t *testing.T) {
	CreateRoleUser(t)
}

func TestGetRole(t *testing.T) {
	role := CreateRoleUser(t)

	getRole, err := testQuery.GetRole(context.Background(), role.ID)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.Equal(t, role.ID, getRole.ID)
	require.Equal(t, role.Name, getRole.Name)
	require.ElementsMatch(t, role.Permission, getRole.Permission)
}

func TestDeleteRow(t *testing.T) {
	role := CreateRoleUser(t)
	err := testQuery.DeleteRole(context.Background(), role.ID)

	require.NoError(t, err)
}
