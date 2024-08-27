package token

import (
	"fmt"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	SymmetricKey []byte
	Paseto       *paseto.V2
}

func NewPasetoMaker(config config.Config) (Maker, error) {

	if len(config.Symmetric) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return &PasetoMaker{
		SymmetricKey: []byte(config.Symmetric),
		Paseto:       paseto.NewV2(),
	}, nil
}

// CreateTokenPaseto implements Maker.
func (p *PasetoMaker) CreateTokenPaseto(id int, permissions []string, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(id, permissions, duration)

	token, err := p.Paseto.Encrypt(p.SymmetricKey, payload, nil)

	if err != nil {
		return "", payload, fmt.Errorf("encrypt fail: %w", err)
	}

	return token, payload, nil
}

// VerifyTokenPaseto implements Maker.
func (p *PasetoMaker) VerifyTokenPaseto(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.Paseto.Decrypt(token, p.SymmetricKey, payload, nil)

	if err != nil {
		return payload, fmt.Errorf(errInvalid)
	}

	if payload.Valid() {
		return payload, fmt.Errorf(errExpired)
	}

	return payload, nil
}
