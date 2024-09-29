package db

import (
	"context"
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func CreateCityUser(t *testing.T) City {
	name := utils.RandomString(10)
	city, err := testQuery.CreateCity(context.Background(), name)

	require.NoError(t, err)
	require.NotZero(t, city.ID)
	require.Equal(t, name, city.Name)

	return city
}

func TestCreateCity(t *testing.T) {
	CreateCityUser(t)
}

func TestGetCityById(t *testing.T) {
	city := CreateCityUser(t)

	getCity, err := testQuery.GetCityById(context.Background(), city.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getCity)
	require.Equal(t, city.ID, getCity.ID)
	require.Equal(t, city.Name, getCity.Name)
}

func TestGetCityByName(t *testing.T) {
	city := CreateCityUser(t)
	getCity, err := testQuery.GetCityByName(context.Background(), city.Name)
	require.NoError(t, err)
	require.NotEmpty(t, getCity)
	require.Equal(t, city.ID, getCity.ID)
	require.Equal(t, city.Name, getCity.Name)
}

func TestUpdateCity(t *testing.T) {
	city := CreateCityUser(t)
	name := utils.RandomString(10)

	arg := UpdateCityParams{
		ID:   city.ID,
		Name: name,
	}

	updateCity, err := testQuery.UpdateCity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateCity)
	require.Equal(t, city.ID, updateCity.ID)
	require.Equal(t, name, updateCity.Name)
}

func TestDeleteCity(t *testing.T) {
	city := CreateCityUser(t)

	err := testQuery.DeleteCityById(context.Background(), int64(city.ID))
	require.NoError(t, err)
}

func TestDeleteManyCity(t *testing.T) {
	city := CreateCityUser(t)
	city1 := CreateCityUser(t)
	city2 := CreateCityUser(t)

	arg := []int64{
		int64(city.ID), int64(city1.ID), int64(city2.ID),
	}

	err := testQuery.DeleteManyCityByIds(context.Background(), arg)
	require.NoError(t, err)
}
