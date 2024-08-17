package token

import "time"

type Maker interface {
	CreateTokenPaseto(id int, duration time.Duration) (string,*Payload,error)
	VerifyTokenPaseto(token string) (*Payload,error)
}