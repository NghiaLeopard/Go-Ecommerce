package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {
	arg := CreateRoleParams{
		Name:       "Basic",
		Permission: []string{},
	}

	role, err := testQuery.CreateRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, role.ID)
	require.Equal(t, arg.Name, role.Name)
	require.Equal(t, arg.Permission, role.Permission)
}
