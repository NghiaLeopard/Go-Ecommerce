package token

import "time"

type Maker interface {
	// Paseto
	// CreateTokenPaseto(id int, permissions []string, duration time.Duration) (string, *Payload, error)
	// VerifyTokenPaseto(token string) (*Payload, error)

	// JWT
	CreateToken(id int, permissions []string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
