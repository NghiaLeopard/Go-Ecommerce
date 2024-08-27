package gmail

import (
	"testing"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestEmailSender(t *testing.T) {
	config, err := config.LoadConfig("../..")

	require.NoError(t, err)

	sender := NewEmailSender(config)

	to := []string{"nghiabeo1605@gmail.com"}
	subject := "Forgot-password"
	text := []byte("Test")

	err = sender.SenderEmail(to, subject, text, nil, nil)

	require.NoError(t, err)

}
