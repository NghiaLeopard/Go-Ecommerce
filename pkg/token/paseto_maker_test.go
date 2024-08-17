package token

import (
	"testing"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateTokenPaseto(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	id := 1
	duration := time.Minute * 60

	paseto,err := NewPasetoMaker([]byte(symmetricKey))

	require.NoError(t,err)

	token,payload1,err := paseto.CreateTokenPaseto(id,duration)

	require.NoError(t,err)
	require.NotEmpty(t,token)
	require.NotEmpty(t,payload1)


	payload2,err := paseto.VerifyTokenPaseto(token)
	t.Log(payload2.IssuedAt,payload2.Expired)

	require.NoError(t,err)
	require.NotEmpty(t,payload2)

	require.Equal(t,id,payload2.Id)
	require.WithinDuration(t,time.Now(),payload2.IssuedAt,time.Second)
	require.WithinDuration(t,time.Now().Add(duration),payload2.Expired,time.Second)
}

func TestTokenExpire(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	id := 1
	duration := time.Minute * 60

	paseto,err := NewPasetoMaker([]byte(symmetricKey))

	require.NoError(t,err)

	token,payload1,err := paseto.CreateTokenPaseto(id,-duration)

	require.NoError(t,err)
	require.NotEmpty(t,token)
	require.NotEmpty(t,payload1)


	_,err = paseto.VerifyTokenPaseto(token)

	require.Error(t,err)
}