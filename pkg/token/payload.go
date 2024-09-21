package token

import (
	"fmt"
	"time"
)

type Payload struct {
	Id          int       `json:"id"`
	Permissions []string  `json:"permissions"`
	IssuedAt    time.Time `json:"issuedAt"`
	Expired     time.Time `json:"expired"`
}

var (
	errInvalid = "token is invalid"
	errExpired = "token is expired"
)

func NewPayload(id int, permissions []string, duration time.Duration) *Payload {
	issuedAt := time.Now()
	expired := issuedAt.Add(duration)

	return &Payload{
		Id:          id,
		Permissions: permissions,
		IssuedAt:    issuedAt,
		Expired:     expired,
	}
}

func (p *Payload) Valid() bool {
	fmt.Println(p.Expired)
	fmt.Println(time.Now())
	return time.Now().After(p.Expired)
}
