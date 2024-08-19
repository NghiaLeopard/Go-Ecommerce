package token

import "time"

type Maker interface {
	CreateTokenPaseto(id int,permissions []string, duration time.Duration) (string,*Payload,error)
	VerifyTokenPaseto(token string) (*Payload,error)
}