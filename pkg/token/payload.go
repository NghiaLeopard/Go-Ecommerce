package token

import (
	"time"
)

type Payload struct {
	Id          int       `json:"id"`
	Permissions []string  `json:"permissions"`
	IssuedAt    time.Time `json:"issuedAt"`
	Expired     int64     `json:"exp"`
}

func NewPayload(id int, permissions []string, duration time.Duration) *Payload {
	issuedAt := time.Now()
	expired := issuedAt.Add(duration)
	expiredTimestamp := expired.Unix()

	return &Payload{
		Id:          id,
		Permissions: permissions,
		IssuedAt:    issuedAt,
		Expired:     expiredTimestamp,
	}
}

func (payload *Payload) Valid() error {
	exp := time.Unix(payload.Expired, 0)
	if time.Now().After(exp) {
		return errExpired
	}
	return nil
}
