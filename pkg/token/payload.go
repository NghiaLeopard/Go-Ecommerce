package token

import "time"

type Payload struct {
	Id       int `json:"id"`
	IssuedAt time.Time `json:"issuedAt"`
	Expired time.Time `json:"expired"`
}

var (
	errInvalid = "token is invalid"
	errExpired = "token is expired"
)

func NewPayload(id int,duration time.Duration) *Payload {
	issuedAt := time.Now()
	expired := issuedAt.Add(duration) 

	return &Payload{
		Id: id,
		IssuedAt: issuedAt,
		Expired: expired,
	}
}

func (p *Payload) Valid() bool {
	return time.Now().After(p.Expired)
}